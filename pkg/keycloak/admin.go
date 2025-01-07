package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AdminService struct {
	AdminUsername string
	AdminPassword string
	Realm         string
	Host          string
}

// NewAdminService initializes a new AdminService
func NewAdminService(adminUsername, adminPassword, realm, host string) *AdminService {
	return &AdminService{
		AdminUsername: adminUsername,
		AdminPassword: adminPassword,
		Realm:         realm,
		Host:          host,
	}
}

// CreateUser creates a new user in Keycloak using the provided email
func (as *AdminService) CreateUser(email string) error {
	url := fmt.Sprintf("%s/auth/admin/realms/%s/users", as.Host, as.Realm)

	userData := map[string]interface{}{
		"email":    email,
		"username": email,
		"enabled":  true,
	}

	jsonData, err := json.Marshal(userData)
	if err != nil {
		return err
	}

	// Get admin access token
	token, err := as.getAdminAccessToken()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create user, status code: %d", resp.StatusCode)
	}

	return nil
}

// getAdminAccessToken retrieves an access token for the admin user
func (as *AdminService) getAdminAccessToken() (string, error) {
	tokenUrl := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", as.Host, as.Realm)

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "admin-cli")
	data.Set("username", as.AdminUsername)
	data.Set("password", as.AdminPassword)

	req, err := http.NewRequest(http.MethodPost, tokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token, status code: %d", resp.StatusCode)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
