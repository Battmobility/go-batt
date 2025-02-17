package batt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchVehicles(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	battClient := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	veh, err := battClient.SearchVehicles(&SearchVehicleRequest{})
	require.NoError(t, err)
	fmt.Println(veh.Vehicles)
}

func TestCreateVBL(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
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
	require.NoError(t, err)
	fmt.Println(vbl) //nolint:forbidigo
}

func TestAddVBLToVehicle(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	veh, err := bc.UpdateVehicleLocation(UpdateVehicleRequest{
		VehicleId: "1THX384",
		AddVehicleLocationRequest: AddVehicleLocationRequest{
			From:                  time.Now().Add(5 * 24 * time.Hour),
			VehicleBaseLocationID: "2eabe147-b394-4c87-ad91-bbd7702f39d2",
		},
	})
	require.NoError(t, err)
	fmt.Println(veh) //nolint:forbidigo
}

func TestGetVehicleGroup(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	vg, err := bc.GetVehicleGroup("batt-all")
	require.NoError(t, err)
	fmt.Println(vg) //nolint:forbidigo
}

func TestCreateIssue(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	bc := NewBattClient("https://booking-staging.battmobility.be/web-api/", "", "https://keycloak-staging.battmobility.be", "batt", os.Getenv("BATT_PASSWORD"))
	issue, err := bc.CreateIssue(CreateIssueRequest{
		Title:     "hoi",
		VehicleId: "2ATZ899",
	})
	require.NoError(t, err)
	fmt.Println(issue) //nolint:forbidigo
}

func TestSearchIssue(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	//create client
	//search for issues from a certain vehicle with a title and statuses CREATED, RESOLVED
	bc := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	issues, err := bc.SearchIssues(SearchIssueRequest{
		VehicleId: "2ATZ899",
		Title:     "carwash",
		Statuses:  []string{"CREATED", "RESOLVED"},
	})
	require.NoError(t, err)
	fmt.Println(issues) //nolint:forbidigo
}

func TestGetNeedsCorrectionBookings(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
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
			require.NoError(t, err)
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

func TestSearchBookings(t *testing.T) {
	t.Parallel()
	if os.Getenv("BATT_PASSWORD") == "" {
		t.Skip("skipping test; BATT_PASSWORD is empty")
	}
	bc := NewBattClient("https://api.battmobility.com/api/web-bff-service/v1/", "", "https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	start := time.Now()
	end := time.Now().Add(24 * time.Hour)
	_, err := bc.SearchBookings(SearchBookingRequest{
		Period: Period{
			Start: &start,
			End:   &end,
		},
	})
	assert.NoError(t, err)
}
