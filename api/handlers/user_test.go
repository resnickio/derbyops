package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/resnickio/derbyops/api/models"
)

func TestGetUsers(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("getUsers() status = %v, want %v", w.Code, http.StatusOK)
	}

	var response map[string][]models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if _, ok := response["users"]; !ok {
		t.Error("Response missing 'users' key")
	}
}

func TestGetUser(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		userID     string
		wantStatus int
	}{
		{
			name:       "existing user",
			userID:     "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existent user",
			userID:     "999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("getUser() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		userID     string
		payload    map[string]string
		wantStatus int
	}{
		{
			name:   "valid update",
			userID: "1",
			payload: map[string]string{
				"firstName": "Updated",
				"lastName":  "User",
				"email":     "updated@example.com",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "non-existent user",
			userID: "999",
			payload: map[string]string{
				"firstName": "Updated",
				"lastName":  "User",
				"email":     "updated@example.com",
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "invalid email",
			userID: "1",
			payload: map[string]string{
				"email": "invalid-email",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("PUT", "/api/v1/users/"+tt.userID, bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("updateUser() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		userID     string
		wantStatus int
	}{
		{
			name:       "existing user",
			userID:     "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existent user",
			userID:     "999",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/api/v1/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("deleteUser() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
} 