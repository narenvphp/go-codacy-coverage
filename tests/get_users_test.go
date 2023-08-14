package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/narenvphp/go-codacy-coverage/controller"
	"github.com/narenvphp/go-codacy-coverage/model"
)

func TestAllUsers(t *testing.T) {
	// Set up a mock HTTP request and response
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Call the handler function
	controller.AllUsers(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var response model.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error parsing response body: %v", err)
	}

	// Check the response data
	expectedStatus := 200
	if response.Status != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, response.Status)
	}

	expectedMessage := "Success"
	if response.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, response.Message)
	}
}
