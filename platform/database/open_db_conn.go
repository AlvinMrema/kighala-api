package database

import database "github.com/AlvinMrema/kighala-api/platform/database/sqlc"

type Queries struct {
	DB *database.Queries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		DB: database.New(db),
	}, nil
}
