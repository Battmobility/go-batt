package batt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearchVehicles(t *testing.T) {
	battClient := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	veh, err := battClient.SearchVehicles(&SearchVehicleRequest{})
	fmt.Print(err)
	fmt.Println(veh.Vehicles)
}

func TestCreateVBL(t *testing.T) {
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	vbl, err := bc.CreateVehicleBaseLocation(VehicleBaseLocation{
		Name: "hoi",
		HomePosition: GpsLocation{
			GpsCoordinateDto: struct {
				Longitude float64 "json:\"longitude\""
				Latitude  float64 "json:\"latitude\""
			}{
				Latitude:  51.2194475,
				Longitude: 4.4024643,
			},
			Address: "Brabantdam 1, 9000 Gent",
		},
	})
	assert.NoError(t, err)
	fmt.Println(vbl)
}

func TestAddVBLToVehicle(t *testing.T) {
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	veh, err := bc.UpdateVehicleLocation(UpdateVehicleRequest{
		VehicleId: "1THX384",
		AddVehicleLocationRequest: AddVehicleLocationRequest{
			From:                  time.Now().Add(5 * 24 * time.Hour),
			VehicleBaseLocationID: "2eabe147-b394-4c87-ad91-bbd7702f39d2",
		},
	})
	assert.NoError(t, err)
	fmt.Println(veh)
}

func TestGetVehicleGroup(t *testing.T) {
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	vg, err := bc.GetVehicleGroup("batt-all")
	assert.NoError(t, err)
	fmt.Println(vg)
}

func TestCreateIssue(t *testing.T) {
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	issue, err := bc.CreateIssue(CreateIssueRequest{
		Title:     "hoi",
		VehicleId: "2ATZ899",
	})
	assert.NoError(t, err)
	fmt.Println(issue)
}

func TestSearchIssue(t *testing.T) {
	//create client
	//search for issues from a certain vehicle with a title and statuses CREATED, RESOLVED
	bc := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	issues, err := bc.SearchIssues(SearchIssueRequest{
		VehicleId: "2ATZ899",
		Title:     "carwash",
		Statuses:  []string{"CREATED", "RESOLVED"},
	})
	assert.NoError(t, err)
	fmt.Println(issues)
}

func TestGetNeedsCorrectionBookings(t *testing.T) {
	needsCorrection := true
	bc := NewBattClient("dummy", "", "dummy", "batt", os.Getenv("BATT_PASSWORD"))
	start := time.Now().Add(-24 * 25 * time.Hour)
	bookings, err := bc.SearchBookings(SearchBookingRequest{
		Period: Period{
			Start: &start,
		},
		NeedsCorrection: &needsCorrection,
	})
	assert.NoError(t, err)
	fmt.Println(len(bookings.Bookings), "problems")
	//for each booking, get the vehicle id
	//and look up the VehicleTelematics
	type mileage struct {
		Mileage   float64   `json:"mileage"`
		Timestamp time.Time `json:"timestamp"`
		DeviceId  string    `json:"deviceId"`
	}
	for _, booking := range bookings.Bookings {
		telematics, err := bc.GetVehicleTelematics(booking.Vehicle.ID)
		assert.NoError(t, err)
		//check if providerId is FLESPI or FLESPI_TWILIO
		if telematics.ProviderId == TELEMATICSPROVIDER_FLESPI || telematics.ProviderId == TELEMATICSPROVIDER_FLESPI_TWILIO {
			//first, parse booking.VehicleUsageDto.StartDate, which has the form 2024-12-29T21:07:57.603996+01:00[Europe/Brussels]
			//as a time.Time object
			sanitizedStartDate := strings.TrimSuffix(booking.VehicleUsageDto.StartDate, "[Europe/Brussels]")
			sanitizedEndDate := strings.TrimSuffix(booking.VehicleUsageDto.EndDate, "[Europe/Brussels]")
			//look up the mileage eg. https://telematics.battmobility.be/flespi/mileage\?deviceId\=5994064\&timestamp\=2024-12-14T22:05:59Z
			//where timestamp is the sanitizedStartDate
			parsedStart, err := time.Parse(time.RFC3339, sanitizedStartDate)
			assert.NoError(t, err)
			parsedEnd, err := time.Parse(time.RFC3339, sanitizedEndDate)
			assert.NoError(t, err)
			duration := parsedEnd.Sub(parsedStart)
			sanitizedStartDate = parsedStart.UTC().Format("2006-01-02T15:04:05Z")
			url := fmt.Sprintf("https://telematics.battmobility.be/flespi/mileage?deviceId=%s&timestamp=%s", telematics.ProviderDeviceId, sanitizedStartDate)
			fmt.Println(url)
			resp, err := http.Get(url)
			assert.NoError(t, err)
			if resp.StatusCode == 200 {
				mileage := &mileage{}
				err = json.NewDecoder(resp.Body).Decode(mileage)
				assert.NoError(t, err)
				delta := booking.VehicleUsageDto.MileageEndValue - mileage.Mileage
				deltaInt := int(delta)
				fmt.Println(delta, "km", fmt.Sprintf("%.1f", duration.Hours()), "h", booking.Vehicle.LicensePlate, booking.VehicleUsageDto.StartDate)
				updateNeedsCorrection := false
				if delta > 0 && delta < 1000 {
					_, err := bc.UpdateBooking(UpdateBookingRequest{
						CorrectedKm:     &deltaInt,
						NeedsCorrection: &updateNeedsCorrection,
					}, booking.ID)
					assert.NoError(t, err)
					fmt.Println("corrected", booking.Vehicle.Name, "by", deltaInt, "km")

				}
			} else {
				respBody, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				fmt.Println("error", resp.StatusCode, string(respBody))
			}
		}
	}
}

func TestGetBookings(t *testing.T) {
	bc := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	eigenBeheerId := "4c7bfd6e-12d6-44f1-8159-88197977d4df"
	start := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	//first get all vehicles in group
	vg, err := bc.GetVehicleGroup(eigenBeheerId)
	assert.NoError(t, err)
	for _, v := range vg.Vehicles {
		//start is 1st of December 2024
		err = CalculateRevenue(bc, start, end, v.ID)
		if err != nil {
			log.Println(err)
		}
	}
}

func CalculateRevenue(bc *BattClient, start, end time.Time, vehicleId string) error {
	exemptedClients := map[string]bool{
		"VlootBeheer":     true,
		"BattMobility NV": true,
	}

	doNotInvoice := false
	bookings, err := bc.SearchBookings(SearchBookingRequest{
		VehicleId:    vehicleId,
		DoNotInvoice: &doNotInvoice,
		EndPeriod: Period{
			Start: &start,
			End:   &end,
		},
	})
	if err != nil {
		return err
	}
	//print user, start date, client, fbp for every booking
	//calculate total hours, total km/hour, total km, total revenue, total number of users
	//exclude bookings from exemptedClients
	totalHours := 0.0
	totalKm := 0
	totalRevenue := 0.0
	totalUsers := 0
	usersCache := map[string]bool{}

	for _, booking := range bookings.Bookings {
		if exemptedClients[booking.Client.Name] {
			continue
		}
		duration := booking.Period.ParsedEnd.Sub(booking.Period.ParsedStart).Hours()
		km := booking.FinishedBookingPrice.Km
		revenue := booking.FinishedBookingPrice.TotalExclVat

		totalHours += duration
		totalKm += km
		totalRevenue += revenue
		//fmt.Println(booking.Period.ParsedStart.Format("2006-01-02"), booking.User.DisplayName)
		if _, ok := usersCache[booking.User.RemoteID]; !ok {
			totalUsers++
			usersCache[booking.User.RemoteID] = true
		}

	}
	fmt.Printf("%s from %v to %v: %.2f hours %d km %.2f km/h %.2f EUR revenue %d users\n", vehicleId, start.Format("2006-01-02"), end.Format("2006-01-02"), totalHours, totalKm, float64(totalKm)/totalHours, totalRevenue, totalUsers)
	return nil
}
