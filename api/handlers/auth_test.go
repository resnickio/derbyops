package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	SetupRoutes(r)
	return r
}

func TestLogin(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		payload    map[string]string
		wantStatus int
	}{
		{
			name: "valid credentials",
			payload: map[string]string{
				"email":    "test@example.com",
				"password": "password123",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "invalid credentials",
			payload: map[string]string{
				"email":    "test@example.com",
				"password": "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "missing email",
			payload: map[string]string{
				"password": "password123",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("login() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		payload    map[string]string
		wantStatus int
	}{
		{
			name: "valid registration",
			payload: map[string]string{
				"email":     "newuser@example.com",
				"password":  "password123",
				"firstName": "New",
				"lastName":  "User",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing required fields",
			payload: map[string]string{
				"email": "test@example.com",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":     "invalid-email",
				"password":  "password123",
				"firstName": "Test",
				"lastName":  "User",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("register() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("logout() status = %v, want %v", w.Code, http.StatusOK)
	}
} 