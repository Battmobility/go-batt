package batt_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/battmobility/go-batt/pkg/batt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ErrDumpReportCsv = errors.New("failed to dump report to csv")
)

func TestGetBookings(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	bc := batt.NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "",
		"https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	eigenBeheerID := "4c7bfd6e-12d6-44f1-8159-88197977d4df"
	vgs, err := bc.GetVehicleGroups("8c2011de-c5fa-4ead-95ef-50c22e5b5b80")
	require.NoError(t, err)
	t.Log(vgs.VehicleGroups)
	start := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	// first get all vehicles in group
	vg, err := bc.GetVehicleGroup(eigenBeheerID)
	require.NoError(t, err)
	vehicleIDs := ""
	for _, v := range vg.Vehicles {
		vehicleIDs += v.ID + ","
	}
	url := "https://backofficetmp.battmobility.be/admin/v1/vehicles/map?sofbattIds=" + vehicleIDs
	res, err := http.Get(url) //nolint:noctx
	require.NoError(t, err)
	defer res.Body.Close()
	backOfficeVehicles := batt.BackOfficeVehicleResponse{}
	err = json.NewDecoder(res.Body).Decode(&backOfficeVehicles)
	require.NoError(t, err)
	reports := []VehicleReport{}
	for _, v := range vg.Vehicles {
		// start is 1st of December 2024
		t.Log("calculating revenue for", v.ID)
		report, err := CalculateRevenue(bc, start, end, v.ID, v.LicensePlate, v.Name)
		if err != nil {
			log.Println(err)
		}
		report.LeasePriceExVat = backOfficeVehicles.Vehicles[v.ID].LeasingMonthlyPriceExVat
		reports = append(reports, *report)
	}
	// now output the reports slice to a csv file
	f, err := os.Create("report.csv")
	require.NoError(t, err)
	defer f.Close()
	_, err = f.WriteString("LicensePlate,Name,Start,End,TotalHours,TotalKm,KmPerHour,KmPerBooking," +
		"TotalBookings,TotalRevenue,TotalUsers,LeasePriceExVat,PnL\n")
	require.NoError(t, err)
	for _, report := range reports {
		_, err = f.WriteString(
			fmt.Sprintf(
				"%s,%s,%s,%s,%.1f,%d,%.2f,%.2f,%d,%.2f,%d,%.2f,%.2f\n",
				report.LicensePlate,
				report.Name,
				report.Start.Format("2006-01-02"),
				report.End.Format("2006-01-02"),
				report.TotalHours,
				report.TotalKm,
				report.KmPerHour,
				report.KmPerBooking,
				report.TotalBookings,
				report.TotalRevenue,
				report.TotalUsers,
				report.LeasePriceExVat,
				report.TotalRevenue-report.LeasePriceExVat,
			))
		assert.NoError(t, err)
	}
}

func CalculateRevenue(bc *batt.Client, start, end time.Time, vehicleID,
	licensePlate, name string) (*VehicleReport, error) {
	exemptedClients := map[string]bool{
		"VlootBeheer":     true,
		"BattMobility NV": true,
	}

	doNotInvoice := false
	bookings, err := bc.SearchBookings(batt.SearchBookingRequest{
		VehicleID:    &vehicleID,
		DoNotInvoice: &doNotInvoice,
		EndPeriod: &batt.Period{
			Start: &start,
			End:   &end,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("searching for bookings failed: %w", err)
	}
	// print user, start date, client, fbp for every booking
	// calculate total hours, total km/hour, total km, total revenue, total number of users
	// exclude bookings from exemptedClients
	totalHours := 0.0
	totalKm := 0
	totalRevenue := 0.0
	totalUsers := 0
	usersCache := map[string]bool{}

	for _, booking := range bookings.Bookings {
		if exemptedClients[booking.Client.Name] {
			continue
		}
		if booking.Vehicle.Owner == booking.User.RemoteID {
			log.Println("skipping owner", booking.User.DisplayName, booking.Vehicle.ID)
			continue
		}
		// fmt.Println(booking.User.DisplayName, booking.Period.ParsedStart.Format("2006-01-02"),
		//     booking.Client.Name, booking.FinishedBookingPrice.TotalExclVat)
		duration := booking.Period.ParsedEnd.Sub(booking.Period.ParsedStart).Hours()
		km := booking.FinishedBookingPrice.Km
		revenue := booking.FinishedBookingPrice.TotalExclVat

		totalHours += duration
		totalKm += km
		totalRevenue += revenue
		// fmt.Println(booking.Period.ParsedStart.Format("2006-01-02"), booking.User.DisplayName)
		if _, ok := usersCache[booking.User.RemoteID]; !ok {
			totalUsers++
			usersCache[booking.User.RemoteID] = true
		}
	}
	return &VehicleReport{
		LicensePlate:  licensePlate,
		Name:          name,
		Start:         start,
		End:           end,
		TotalHours:    totalHours,
		TotalKm:       totalKm,
		KmPerBooking:  float64(totalKm) / float64(len(bookings.Bookings)),
		KmPerHour:     float64(totalKm) / totalHours,
		TotalRevenue:  totalRevenue,
		TotalBookings: len(bookings.Bookings),
		TotalUsers:    totalUsers,
	}, nil
}

type VehicleReport struct {
	LicensePlate    string
	Name            string
	Start           time.Time
	End             time.Time
	TotalHours      float64
	HoursPerBooking float64
	TotalKm         int
	KmPerBooking    float64
	KmPerHour       float64
	TotalRevenue    float64
	TotalUsers      int
	TotalBookings   int
	LeasePriceExVat float64
	PnL             float64
}

func DumpReportCsv(filename string, reports []VehicleReport) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDumpReportCsv, err)
	}
	defer f.Close()
	_, err = f.WriteString("LicensePlate,Name,Start,End,TotalHours,TotalKm,KmPerHour,KmPerBooking," +
		"TotalBookings,TotalRevenue,TotalUsers,LeasePriceExVat,PnL\n")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDumpReportCsv, err)
	}
	for _, report := range reports {
		_, err = f.WriteString(
			fmt.Sprintf(
				"%s,%s,%s,%s,%.1f,%d,%.2f,%.2f,%d,%.2f,%d,%.2f,%.2f\n",
				report.LicensePlate,
				report.Name,
				report.Start.Format("2006-01-02"),
				report.End.Format("2006-01-02"),
				report.TotalHours,
				report.TotalKm,
				report.KmPerHour,
				report.KmPerBooking,
				report.TotalBookings,
				report.TotalRevenue,
				report.TotalUsers,
				report.LeasePriceExVat,
				report.TotalRevenue-report.LeasePriceExVat,
			))
		if err != nil {
			return fmt.Errorf("%w: %w", ErrDumpReportCsv, err)
		}
	}
	return nil
}
