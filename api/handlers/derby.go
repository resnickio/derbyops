package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/resnickio/derbyops/api/models"
)

// getDerbies returns a list of all derbies
func getDerbies(c *gin.Context) {
	var derbies []models.Derby
	// TODO: Implement database query
	c.JSON(http.StatusOK, gin.H{"derbies": derbies})
}

// getDerby returns a single derby by ID
func getDerby(c *gin.Context) {
	id := c.Param("id")
	var derby models.Derby
	// TODO: Implement database query
	if derby.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Derby not found"})
		return
	}
	c.JSON(http.StatusOK, derby.ToResponse())
}

// createDerby creates a new derby
func createDerby(c *gin.Context) {
	var input struct {
		Name        string    `json:"name" binding:"required"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		StartDate   time.Time `json:"startDate"`
		EndDate     time.Time `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Get user ID from context
	userID := uint(1) // Placeholder

	derby := models.Derby{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		CreatedBy:   userID,
	}

	// TODO: Save to database

	c.JSON(http.StatusCreated, derby.ToResponse())
}

// updateDerby updates a derby's information
func updateDerby(c *gin.Context) {
	id := c.Param("id")
	var derby models.Derby
	// TODO: Implement database query
	if derby.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Derby not found"})
		return
	}

	var input struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		StartDate   time.Time `json:"startDate"`
		EndDate     time.Time `json:"endDate"`
		Status      string    `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update derby in database
	derby.Name = input.Name
	derby.Description = input.Description
	derby.Location = input.Location
	derby.StartDate = input.StartDate
	derby.EndDate = input.EndDate
	derby.Status = input.Status

	c.JSON(http.StatusOK, derby.ToResponse())
}

// deleteDerby deletes a derby
func deleteDerby(c *gin.Context) {
	id := c.Param("id")
	// TODO: Implement database deletion
	c.JSON(http.StatusOK, gin.H{"message": "Derby deleted successfully"})
} 