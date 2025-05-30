package jsonplaceholder

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type JSONPlaceholderClientInterface interface {
	SyncUsers() ([]User, error)
}

type JSONPlaceholderClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewJSONPlaceholderClient(baseURL string) JSONPlaceholderClientInterface {
	return &JSONPlaceholderClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *JSONPlaceholderClient) SyncUsers() ([]User, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch users")
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
