package models

import (
	"time"
)

// Derby represents a roller derby event or league
type Derby struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Status      string    `json:"status" gorm:"default:'scheduled'"` // scheduled, in_progress, completed, cancelled
	CreatedBy   uint      `json:"createdBy" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// DerbyResponse represents the derby data sent in API responses
type DerbyResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Status      string    `json:"status"`
	CreatedBy   uint      `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ToResponse converts a Derby to a DerbyResponse
func (d *Derby) ToResponse() *DerbyResponse {
	return &DerbyResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Location:    d.Location,
		StartDate:   d.StartDate,
		EndDate:     d.EndDate,
		Status:      d.Status,
		CreatedBy:   d.CreatedBy,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
} 