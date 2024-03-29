package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sql.DB, error) {
	// Define database connection settings.
	// maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	// maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	// maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Define database connection for PostgreSQL.
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	// Set database connection settings.
	// db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	// db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	// db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
