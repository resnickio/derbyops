package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes
func SetupRoutes(r *gin.Engine) {
	// Health check endpoint
	r.GET("/health", healthCheck)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", login)
			auth.POST("/register", register)
			auth.POST("/logout", logout)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(authMiddleware())
		{
			// User routes
			protected.GET("/users", getUsers)
			protected.GET("/users/:id", getUser)
			protected.PUT("/users/:id", updateUser)
			protected.DELETE("/users/:id", deleteUser)

			// Derby routes
			protected.GET("/derbies", getDerbies)
			protected.POST("/derbies", createDerby)
			protected.GET("/derbies/:id", getDerby)
			protected.PUT("/derbies/:id", updateDerby)
			protected.DELETE("/derbies/:id", deleteDerby)
		}
	}
}

// healthCheck handles the health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "DerbyOps API is healthy",
	})
}

// authMiddleware ensures the request is authenticated
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT token validation
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		// TODO: Validate token and set user context
		c.Next()
	}
} 