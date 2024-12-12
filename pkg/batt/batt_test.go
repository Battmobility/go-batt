package batt

import (
	"fmt"
	"os"
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
