package batt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BOOKING   = "BOOKING"
	BLOCKED   = "BLOCKED"
	ACTIVE    = "ACTIVE"
	BattOrgID = "8c2011de-c5fa-4ead-95ef-50c22e5b5b80"
	timeout   = 10 * time.Second
)

type Client struct {
	SofBattBaseURL string
	AuthURL        string
	BackOfficeHost string
	APIEndpoint    string
	Token          *string
	Username       string
	Password       string
}

var (
	ErrBadResponse    = errors.New("bad response")
	ErrUserNotFound   = errors.New("user not found")
	ErrMarshalRequest = errors.New("failed to marshal request")
	ErrDecodeResponse = errors.New("failed to decode response")
)

const layout = "2006-01-02T15:04:05Z[UTC]"

func NewBattClient(base, backoffice, auth, username, password string) *Client {
	return &Client{
		SofBattBaseURL: base,
		BackOfficeHost: backoffice,
		AuthURL:        auth,
		Username:       username,
		Password:       password,
	}
}

func (c *Client) refreshToken(ctx context.Context) error {
	data := url.Values{}
	data.Set("client_id", "web")
	data.Set("username", c.Username)
	data.Set("password", c.Password)
	data.Set("grant_type", "password")

	req, err := http.NewRequest(http.MethodPost, c.AuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%w: %s", ErrBadResponse, string(body))
	}
	defer resp.Body.Close()
	token := &Token{}
	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	c.Token = &token.AccessToken
	return nil
}

func (c *Client) SearchAvailabilities(start, end *time.Time, vehicleIDs []string) (*AvailabilityEventResponse, error) {
	requestBody := SearchAvailabilityRequest{
		Period: Period{
			Start: start,
			End:   end,
		},
		VehicleIDs: vehicleIDs,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "availability/v1/availability-events", jsonData)
	if err != nil {
		return nil, err
	}
	res := &AvailabilityEventResponse{}
	err = json.NewDecoder(resp).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
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
	return res, nil
}

func (c *Client) SearchBookings(request SearchBookingRequest) (*SearchBookingResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "booking/v1/bookings/searches", jsonData)
	if err != nil {
		return nil, err
	}
	res := &SearchBookingResponse{}
	err = json.NewDecoder(resp).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal booking response: %w", err)
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

func (c *Client) GetBatteryStatus(id string) (*BatteryStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodGet, c.SofBattBaseURL, "telematics/v1/batteryStatus/"+id, nil)
	if err != nil {
		return nil, err
	}
	res := &BatteryStatus{}
	err = json.NewDecoder(resp).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return res, nil
}

func (c *Client) RefreshLocation(id string) (*GpsLocation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "telematics/v1/location/"+id, nil)
	if err != nil {
		return nil, err
	}
	res := &GpsLocation{}
	err = json.NewDecoder(resp).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return res, nil
}

func (c *Client) SearchVehicles(req *SearchVehicleRequest) (*SearchVehicleResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "booking/v1/vehicles/searches", jsonData)
	if err != nil {
		return nil, err
	}
	respBody := SearchVehicleResponse{}
	err = json.NewDecoder(resp).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return &respBody, nil
}

func (c *Client) GetVehicle(vehicleID string) (*Vehicle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodGet, c.SofBattBaseURL, "booking/v1/vehicles/"+vehicleID, nil)
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

func (c *Client) GetVehicleGroups(organizationID string) (*VehicleGroupsPage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodGet, c.SofBattBaseURL,
		"vehicle-group/v1/vehicle-groups?organizationId="+organizationID, nil)
	if err != nil {
		return nil, err
	}
	result := &VehicleGroupsPage{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) GetVehicleGroup(id string) (*VehicleGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodGet, c.SofBattBaseURL, "vehicle-group/v1/vehicle-groups/"+id, nil)
	if err != nil {
		return nil, err
	}
	result := &VehicleGroup{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) CreateIssue(req CreateIssueRequest) (*Issue, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "issue/v1/issues", jsonData)
	if err != nil {
		return nil, err
	}
	result := &Issue{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) GetVehicleTelematics(id string) (*VehicleTelematics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodGet, c.SofBattBaseURL, "telematics/v1/devices/"+id, nil)
	if err != nil {
		return nil, err
	}
	result := &VehicleTelematics{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) SearchIssues(req SearchIssueRequest) (*IssueResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.SofBattBaseURL, "issue/v1/issues/searches", jsonData)
	if err != nil {
		return nil, err
	}
	result := &IssueResponse{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}
func (c *Client) UpdateBooking(req UpdateBookingRequest, id string) (*Booking, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPut, c.SofBattBaseURL, "booking/v1/bookings/"+id, jsonData)
	if err != nil {
		return nil, err
	}
	result := &Booking{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	return result, nil
}

func (c *Client) SearchBackOfficeUser(remoteID string) (*BackOfficeUser, error) {
	req := SearchBackOfficeUsersRequest{
		SofBattRemoteID: strings.ToUpper(remoteID),
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshalRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.request(ctx, http.MethodPost, c.BackOfficeHost, "/v1/users/search", jsonData)
	if err != nil {
		return nil, err
	}
	result := &BackOfficeUserResponse{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodeResponse, err)
	}
	if len(result.Users) == 0 {
		return nil, fmt.Errorf("%w: sofbatt id %s", ErrUserNotFound, remoteID)
	}
	return &result.Users[0], nil
}

func (c *Client) request(ctx context.Context, method, host, endpoint string, jsonData []byte) (io.ReadCloser, error) {
	if c.Token == nil {
		err := c.refreshToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+*c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		msgB, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%w: %s msg %s", ErrBadResponse, resp.Status, string(msgB))
	}
	return resp.Body, nil
}
