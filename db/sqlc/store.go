package db

import "database/sql"

// Store provides all functions to execute db queries
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates an instance of a store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
