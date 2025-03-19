// Package database provides database connection and migration functionality.
package database

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/resnickio/derbyops/api/logger"
)

// RunMigrations executes database migrations using golang-migrate.
// It applies all pending migrations in the specified direction.
func RunMigrations(direction string) error {
	// Get database instance
	db, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Create postgres driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Create migrate instance
	migrationsPath := fmt.Sprintf("file://%s/migrations", os.Getenv("APP_ROOT"))
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Run migrations
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run migrations up: %w", err)
		}
		logger.Info("Successfully ran migrations up", nil)
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run migrations down: %w", err)
		}
		logger.Info("Successfully ran migrations down", nil)
	default:
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	return nil
}

// MigrateToVersion migrates the database to a specific version.
// This is useful for rolling back to a specific version or jumping to a version.
func MigrateToVersion(version uint) error {
	// Get database instance
	db, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Create postgres driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Create migrate instance
	migrationsPath := fmt.Sprintf("file://%s/migrations", os.Getenv("APP_ROOT"))
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Migrate to specific version
	if err := m.Migrate(version); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate to version %d: %w", version, err)
	}

	logger.Info("Successfully migrated to version", map[string]interface{}{
		"version": version,
	})

	return nil
} 