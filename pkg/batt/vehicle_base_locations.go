package batt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) CreateVehicleBaseLocation(request VehicleBaseLocation) (*VehicleBaseLocation, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "vehicle/v1/vehiclebaselocations", reqBytes)
	if err != nil {
		return nil, err
	}
	result := &VehicleBaseLocation{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) UpdateVehicleLocation(request UpdateVehicleRequest) (*Vehicle, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPut, c.SofBattBaseURL,
		"booking/v1/vehicles/"+request.VehicleID, reqBytes)
	if err != nil {
		return nil, err
	}
	result := &Vehicle{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

type VehicleBaseLocation struct {
	ID           string      `json:"id,omitempty"`
	Name         string      `json:"name"`
	HomePosition GpsLocation `json:"homePosition"`
	Borough      string      `json:"borough"`
	Status       string      `json:"status"`
	Memo         string      `json:"memo"`
	ADType       string      `json:"adType"`
}

type UpdateVehicleRequest struct {
	VehicleID                 string                    `json:"id,omitempty"`
	AddVehicleLocationRequest AddVehicleLocationRequest `json:"addVehicleLocationRequest"`
}

type AddVehicleLocationRequest struct {
	From                  time.Time `json:"from"`
	VehicleBaseLocationID string    `json:"vehicleBaseLocationId"`
}
