package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resnickio/derbyops/api/models"
)

// getUsers returns a list of all users
func getUsers(c *gin.Context) {
	var users []models.User
	// TODO: Implement database query
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// getUser returns a single user by ID
func getUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	// TODO: Implement database query
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user.ToResponse())
}

// updateUser updates a user's information
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	// TODO: Implement database query
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update user in database
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email

	c.JSON(http.StatusOK, user.ToResponse())
}

// deleteUser deletes a user
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Implement database deletion
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
} 