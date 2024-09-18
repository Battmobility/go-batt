package batt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *BattClient) CreateVehicleBaseLocation(request VehicleBaseLocation) (result *VehicleBaseLocation, err error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPost, c.SofBattBaseUrl, "vehicle/v1/vehiclebaselocations", reqBytes)
	if err != nil {
		return nil, err
	}
	result = &VehicleBaseLocation{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *BattClient) UpdateVehicleLocation(request UpdateVehicleRequest) (result *Vehicle, err error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPut, c.SofBattBaseUrl, fmt.Sprintf("booking/v1/vehicles/%s", request.VehicleId), reqBytes)
	if err != nil {
		return nil, err
	}
	result = &Vehicle{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type VehicleBaseLocation struct {
	ID           string      `json:"id,omitempty"`
	Name         string      `json:"name"`
	HomePosition GpsLocation `json:"homePosition"`
}

type UpdateVehicleRequest struct {
	VehicleId                 string                    `json:"id,omitempty"`
	AddVehicleLocationRequest AddVehicleLocationRequest `json:"addVehicleLocationRequest"`
}

type AddVehicleLocationRequest struct {
	From                  time.Time `json:"from"`
	VehicleBaseLocationID string    `json:"vehicleBaseLocationId"`
}
