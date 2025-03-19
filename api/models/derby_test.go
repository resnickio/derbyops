package models

import (
	"testing"
	"time"
)

func TestDerby_ToResponse(t *testing.T) {
	now := time.Now()
	d := &Derby{
		ID:          1,
		Name:        "Test Derby",
		Description: "Test Description",
		Location:    "Test Location",
		StartDate:   now,
		EndDate:     now.Add(24 * time.Hour),
		Status:      "scheduled",
		CreatedBy:   1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	resp := d.ToResponse()
	if resp.ID != d.ID {
		t.Errorf("DerbyResponse.ID = %v, want %v", resp.ID, d.ID)
	}
	if resp.Name != d.Name {
		t.Errorf("DerbyResponse.Name = %v, want %v", resp.Name, d.Name)
	}
	if resp.Description != d.Description {
		t.Errorf("DerbyResponse.Description = %v, want %v", resp.Description, d.Description)
	}
	if resp.Location != d.Location {
		t.Errorf("DerbyResponse.Location = %v, want %v", resp.Location, d.Location)
	}
	if !resp.StartDate.Equal(d.StartDate) {
		t.Errorf("DerbyResponse.StartDate = %v, want %v", resp.StartDate, d.StartDate)
	}
	if !resp.EndDate.Equal(d.EndDate) {
		t.Errorf("DerbyResponse.EndDate = %v, want %v", resp.EndDate, d.EndDate)
	}
	if resp.Status != d.Status {
		t.Errorf("DerbyResponse.Status = %v, want %v", resp.Status, d.Status)
	}
	if resp.CreatedBy != d.CreatedBy {
		t.Errorf("DerbyResponse.CreatedBy = %v, want %v", resp.CreatedBy, d.CreatedBy)
	}
	if !resp.CreatedAt.Equal(d.CreatedAt) {
		t.Errorf("DerbyResponse.CreatedAt = %v, want %v", resp.CreatedAt, d.CreatedAt)
	}
	if !resp.UpdatedAt.Equal(d.UpdatedAt) {
		t.Errorf("DerbyResponse.UpdatedAt = %v, want %v", resp.UpdatedAt, d.UpdatedAt)
	}
} 