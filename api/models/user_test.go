package models

import (
	"testing"
)

func TestUser_SetPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "validPassword123",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			err := u.SetPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.SetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Password == "" {
				t.Error("User.SetPassword() password was not set")
			}
		})
	}
}

func TestUser_CheckPassword(t *testing.T) {
	u := &User{}
	password := "testPassword123"
	err := u.SetPassword(password)
	if err != nil {
		t.Fatalf("Failed to set password: %v", err)
	}

	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "correct password",
			password: password,
			want:     true,
		},
		{
			name:     "incorrect password",
			password: "wrongPassword",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := u.CheckPassword(tt.password); got != tt.want {
				t.Errorf("User.CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ToResponse(t *testing.T) {
	u := &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Role:      "user",
	}

	resp := u.ToResponse()
	if resp.ID != u.ID {
		t.Errorf("UserResponse.ID = %v, want %v", resp.ID, u.ID)
	}
	if resp.Email != u.Email {
		t.Errorf("UserResponse.Email = %v, want %v", resp.Email, u.Email)
	}
	if resp.FirstName != u.FirstName {
		t.Errorf("UserResponse.FirstName = %v, want %v", resp.FirstName, u.FirstName)
	}
	if resp.LastName != u.LastName {
		t.Errorf("UserResponse.LastName = %v, want %v", resp.LastName, u.LastName)
	}
	if resp.Role != u.Role {
		t.Errorf("UserResponse.Role = %v, want %v", resp.Role, u.Role)
	}
} 