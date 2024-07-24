package database

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
)

//go:embed schemas.sql
var schemaSQL string

func InitDB(url string) *sql.DB{
	// Open database
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Create tables
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		log.Fatal("Error creating tables: ", err)
	}
	return db
}