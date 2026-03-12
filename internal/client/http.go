package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HTTPClient struct {
	baseURL string
	client  *http.Client
}

type SubmitTaskRequest struct {
	Command string `json:"command"`
}

type SubmitTaskResponse struct {
	Status  string `json:"status"`
	Command string `json:"command"`
}

func NewHTTPClient(baseURL string) (*HTTPClient, error) {
	normalizedURL, err := normalizeBaseURL(baseURL)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{
		baseURL: normalizedURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

func normalizeBaseURL(raw string) (string, error) {
	addr := strings.TrimSpace(raw)
	if addr == "" {
		return "", fmt.Errorf("http address cannot be empty")
	}

	if strings.HasPrefix(addr, ":") {
		addr = "127.0.0.1" + addr
	}

	if !strings.Contains(addr, "://") {
		addr = "http://" + addr
	}

	parsed, err := url.Parse(addr)
	if err != nil {
		return "", fmt.Errorf("invalid http address %q: %w", raw, err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("invalid http address %q: expected host:port or URL", raw)
	}

	return strings.TrimRight(parsed.String(), "/"), nil
}

func (c *HTTPClient) SubmitTask(ctx context.Context, command string) (*SubmitTaskResponse, error) {
	body, err := json.Marshal(SubmitTaskRequest{Command: command})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/tasks",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var out SubmitTaskResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *HTTPClient) HealthCheck(ctx context.Context) error {
	body := []byte(`{"command":"ping"}`)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/tasks",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}
