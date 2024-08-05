package batt

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchVehicles(t *testing.T) {
	battClient := NewBattClient("https://api.battmobility.com", "batt", os.Getenv("BATT_PASSWORD"))
	veh, err := battClient.SearchVehicles(&SearchVehicleRequest{})
	assert.NoError(t, err)
	fmt.Println(veh.Vehicles)
}
