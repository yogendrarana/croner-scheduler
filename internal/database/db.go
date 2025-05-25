package database

import (
	"log"
	"sync"
	"time"

	"croner-scheduler/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	db   *sqlx.DB
	once sync.Once
)

// NewDB creates a new database connection using sqlx
func NewDB(cfg *config.Config) *sqlx.DB {
	once.Do(func() {
		var err error

		// Connect to database
		db, err = sqlx.Connect("postgres", cfg.DBUrl)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Configure connection pool
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		log.Println("Database connection established.")
	})

	return db
}

// CloseDB closes the database connection
func CloseDB() {
	if db == nil {
		return
	}

	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}

	log.Printf("Database connection closed.")
}

// GetDB returns the underlying *sqlx.DB instance
func GetDB() *sqlx.DB {
	return db
}
