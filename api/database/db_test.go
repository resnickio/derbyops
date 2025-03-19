package database

import (
	"os"
	"testing"

	"github.com/resnickio/derbyops/api/models"
)

func TestInitDB(t *testing.T) {
	// Set test environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "derbyops_test")
	os.Setenv("DB_PORT", "5432")

	// Initialize database
	InitDB()

	// Test database connection
	if DB == nil {
		t.Error("Database connection is nil")
	}

	// Test auto-migration
	var user models.User
	var derby models.Derby

	// Check if tables exist
	err := DB.Migrator().HasTable(&user)
	if err != nil {
		t.Errorf("Failed to check User table: %v", err)
	}
	if !DB.Migrator().HasTable(&user) {
		t.Error("User table was not created")
	}

	err = DB.Migrator().HasTable(&derby)
	if err != nil {
		t.Errorf("Failed to check Derby table: %v", err)
	}
	if !DB.Migrator().HasTable(&derby) {
		t.Error("Derby table was not created")
	}

	// Test database operations
	testUser := models.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	err = testUser.SetPassword("password123")
	if err != nil {
		t.Fatalf("Failed to set password: %v", err)
	}

	// Create test user
	err = DB.Create(&testUser).Error
	if err != nil {
		t.Errorf("Failed to create test user: %v", err)
	}

	// Read test user
	var retrievedUser models.User
	err = DB.First(&retrievedUser, testUser.ID).Error
	if err != nil {
		t.Errorf("Failed to retrieve test user: %v", err)
	}

	// Verify user data
	if retrievedUser.Email != testUser.Email {
		t.Errorf("Retrieved user email = %v, want %v", retrievedUser.Email, testUser.Email)
	}
	if retrievedUser.FirstName != testUser.FirstName {
		t.Errorf("Retrieved user firstName = %v, want %v", retrievedUser.FirstName, testUser.FirstName)
	}
	if retrievedUser.LastName != testUser.LastName {
		t.Errorf("Retrieved user lastName = %v, want %v", retrievedUser.LastName, testUser.LastName)
	}

	// Clean up
	DB.Unscoped().Delete(&testUser)
}

func TestDBConnection(t *testing.T) {
	if DB == nil {
		t.Error("Database connection is nil")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Errorf("Failed to ping database: %v", err)
	}
} 