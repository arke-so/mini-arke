package testutil

import (
	"gorm.io/gorm/logger"
	"testing"

	"github.com/arke-so/mini-arke/internal/database"
	"gorm.io/gorm"
)

const TestDatabaseURL = "postgres://developer:devpassword@localhost:5433/mini_arke_test_db?sslmode=disable"

// SetupTestDB creates a fresh database connection and runs migrations
func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := database.NewPostgresDBFromURL(TestDatabaseURL)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	db.Logger = logger.Default.LogMode(logger.Silent)

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return db
}

func CleanupTestDB(t *testing.T, db *gorm.DB) {
	t.Helper()

	db.Exec("DROP TABLE IF EXISTS order_items CASCADE")
	db.Exec("DROP TABLE IF EXISTS orders CASCADE")
	db.Exec("DROP TABLE IF EXISTS products CASCADE")
	db.Exec("DROP TABLE IF EXISTS customers CASCADE")
}
