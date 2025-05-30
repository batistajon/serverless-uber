package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type UberService struct {
	BaseUrl string
}

func NewUberService(clientID string, clientSecret string) (*UberService, error) {
	return &UberService{
		BaseUrl: "https://sandbox-api.uber.com",
	}, nil
}

func (us *UberService) GetAuthToken(clientID string, clientSecret string) (string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	writer.WriteField("client_id", clientID)
	writer.WriteField("client_secret", clientSecret)
	writer.WriteField("grant_type", "client_credentials")
	writer.WriteField("scope", "guests.trips")
	writer.Close()

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://auth.uber.com/oauth/v2/token", &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", body)
	}

	// Extract the access_token from response
	var result struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}
