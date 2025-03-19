package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/resnickio/derbyops/api/models"
)

func TestGetDerbies(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/api/v1/derbies", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("getDerbies() status = %v, want %v", w.Code, http.StatusOK)
	}

	var response map[string][]models.DerbyResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if _, ok := response["derbies"]; !ok {
		t.Error("Response missing 'derbies' key")
	}
}

func TestGetDerby(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		derbyID    string
		wantStatus int
	}{
		{
			name:       "existing derby",
			derbyID:    "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existent derby",
			derbyID:    "999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/derbies/"+tt.derbyID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("getDerby() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestCreateDerby(t *testing.T) {
	router := setupTestRouter()

	now := time.Now()
	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "valid derby",
			payload: map[string]interface{}{
				"name":        "Test Derby",
				"description": "Test Description",
				"location":    "Test Location",
				"startDate":   now,
				"endDate":     now.Add(24 * time.Hour),
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing required fields",
			payload: map[string]interface{}{
				"description": "Test Description",
				"location":    "Test Location",
				"startDate":   now,
				"endDate":     now.Add(24 * time.Hour),
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "invalid dates",
			payload: map[string]interface{}{
				"name":        "Test Derby",
				"description": "Test Description",
				"location":    "Test Location",
				"startDate":   now.Add(24 * time.Hour),
				"endDate":     now,
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/derbies", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("createDerby() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestUpdateDerby(t *testing.T) {
	router := setupTestRouter()

	now := time.Now()
	tests := []struct {
		name       string
		derbyID    string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name:    "valid update",
			derbyID: "1",
			payload: map[string]interface{}{
				"name":        "Updated Derby",
				"description": "Updated Description",
				"location":    "Updated Location",
				"startDate":   now,
				"endDate":     now.Add(24 * time.Hour),
				"status":      "scheduled",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:    "non-existent derby",
			derbyID: "999",
			payload: map[string]interface{}{
				"name": "Updated Derby",
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:    "invalid status",
			derbyID: "1",
			payload: map[string]interface{}{
				"status": "invalid_status",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("PUT", "/api/v1/derbies/"+tt.derbyID, bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("updateDerby() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestDeleteDerby(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		derbyID    string
		wantStatus int
	}{
		{
			name:       "existing derby",
			derbyID:    "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existent derby",
			derbyID:    "999",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/api/v1/derbies/"+tt.derbyID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("deleteDerby() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
} 