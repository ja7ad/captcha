package captcha

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockResponse struct {
	Message string `json:"message"`
}

func TestClient_Request_SuccessWithResponse(t *testing.T) {
	expected := mockResponse{Message: "hello"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "123", r.URL.Query().Get("id"))

		_ = json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	c := newClient()
	ctx := context.Background()

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"id": "123",
	}

	var result mockResponse
	err := c.Request(ctx, server.URL, "GET", nil, &result, headers, params)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestClient_Request_SuccessWithBody(t *testing.T) {
	expected := mockResponse{Message: "received"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		assert.Contains(t, string(body), `"input":"hi"`)

		_ = json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	c := newClient()
	ctx := context.Background()

	body := map[string]string{"input": "hi"}
	var result mockResponse
	err := c.Request(ctx, server.URL, "POST", body, &result, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestClient_Request_Non2xxResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newClient()
	ctx := context.Background()

	err := c.Request(ctx, server.URL, "GET", nil, nil, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "non-2xx response")
}

func TestClient_Request_InvalidJSONBody(t *testing.T) {
	c := newClient()
	ctx := context.Background()

	badBody := map[string]any{
		"invalid": make(chan int),
	}

	err := c.Request(ctx, "http://example.com", "POST", badBody, nil, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "json: unsupported type")
}

func TestClient_Request_BadURL(t *testing.T) {
	c := newClient()
	ctx := context.Background()

	err := c.Request(ctx, ":", "GET", nil, nil, nil, nil)
	assert.Error(t, err)
}
