package batt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BOOKING = "BOOKING"
	BLOCKED = "BLOCKED"
	ACTIVE  = "ACTIVE"
)

type BattClient struct {
	SofBattBaseUrl string
	AuthBaseUrl    string
	BackOfficeHost string
	ApiEndpoint    string
	Token          *string
	Username       string
	Password       string
}

const layout = "2006-01-02T15:04:05-07:00[Europe/Brussels]"

func NewBattClient(base, auth, username, password string) (result *BattClient) {
	return &BattClient{
		SofBattBaseUrl: base,
		AuthBaseUrl:    auth,
		Username:       username,
		Password:       password,
	}
}

func (c *BattClient) refreshToken() error {
	fmt.Println("refreshing token")
	data := url.Values{}
	data.Set("client_id", "web")
	data.Set("username", c.Username)
	data.Set("password", c.Password)
	data.Set("grant_type", "password")

	endpoint := "/auth/realms/Battmobiel/protocol/openid-connect/token"
	url := fmt.Sprintf("%s%s", c.AuthBaseUrl, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	token := &Token{}
	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return err
	}
	c.Token = &token.AccessToken
	return nil
}

func (c *BattClient) SearchAvailabilities(start, end *time.Time, vehicleIds []string) (res *AvailabilityEventResponse, err error) {
	requestBody := SearchAvailabilityRequest{
		Period: Period{
			Start: start,
			End:   end,
		},
		VehicleIds: vehicleIds,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPost, c.SofBattBaseUrl, "availability/v1/availability-events", jsonData)
	if err != nil {
		return nil, err
	}
	res = &AvailabilityEventResponse{}
	err = json.NewDecoder(resp).Decode(res)
	for id, e := range res.Events {
		for i, nav := range e.NonAvailabilities {
			st, _ := time.Parse(layout, nav.Period.Start)
			end, _ := time.Parse(layout, nav.Period.End)
			nav.Period.ParsedStart = st
			nav.Period.ParsedEnd = end
			e.NonAvailabilities[i] = nav
			res.Events[id] = e
		}
	}
	return res, err
}

func (c *BattClient) SearchBookings(request SearchBookingRequest) (res *SearchBookingResponse, err error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPost, c.SofBattBaseUrl, "booking/v1/bookings/searches", jsonData)
	if err != nil {
		return nil, err
	}
	res = &SearchBookingResponse{}
	err = json.NewDecoder(resp).Decode(res)
	if err != nil {
		return nil, err
	}

	for i, b := range res.Bookings {
		s, _ := time.Parse(layout, b.Period.Start)
		e, _ := time.Parse(layout, b.Period.End)
		b.Period.ParsedStart = s
		b.Period.ParsedEnd = e
		res.Bookings[i] = b
	}
	return res, nil
}

func (c *BattClient) GetBatteryStatus(id string) (res *BatteryStatus, err error) {
	resp, err := c.request(http.MethodGet, c.SofBattBaseUrl, fmt.Sprintf("telematics/v1/batteryStatus/%s", id), nil)
	if err != nil {
		return nil, err
	}
	res = &BatteryStatus{}
	err = json.NewDecoder(resp).Decode(res)
	return res, err
}

func (c *BattClient) RefreshLocation(id string) (res *GpsLocation, err error) {
	resp, err := c.request(http.MethodPost, c.SofBattBaseUrl, fmt.Sprintf("telematics/v1/location/%s", id), nil)
	if err != nil {
		return nil, err
	}
	res = &GpsLocation{}
	err = json.NewDecoder(resp).Decode(res)
	return res, err
}

func (c *BattClient) SearchVehicles(req *SearchVehicleRequest) (res *SearchVehicleResponse, err error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPost, c.SofBattBaseUrl, "booking/v1/vehicles/searches", jsonData)
	if err != nil {
		return nil, err
	}
	respBody := SearchVehicleResponse{}
	err = json.NewDecoder(resp).Decode(&respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

func (c *BattClient) GetVehicle(vehicleId string) (res *Vehicle, err error) {
	resp, err := c.request(http.MethodGet, c.SofBattBaseUrl, fmt.Sprintf("booking/v1/vehicles/%s", vehicleId), nil)
	if err != nil {
		return nil, err
	}
	result := &Vehicle{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *BattClient) SearchBackOfficeUser(remoteId string) (res *BackOfficeUser, err error) {
	req := SearchBackOfficeUsersRequest{
		SofBattRemoteId: strings.ToUpper(remoteId),
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request(http.MethodPost, c.BackOfficeHost, "/v1/users/search", jsonData)
	if err != nil {
		return nil, err
	}
	result := &BackOfficeUserResponse{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, err
	}
	if len(result.Users) == 0 {
		return nil, fmt.Errorf("non user found for sofbatt id %s", remoteId)
	}
	return &result.Users[0], nil
}

func (c *BattClient) request(method, host, endpoint string, jsonData []byte) (result io.ReadCloser, err error) {
	if c.Token == nil {
		err = c.refreshToken()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *c.Token))
	req.Header.Set("Content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		msgB, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad response code: %s msg %s", resp.Status, string(msgB))
	}
	return resp.Body, nil
}
