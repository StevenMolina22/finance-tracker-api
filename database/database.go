package database

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

//go:embed schemas/schemas.sql
var schemaSQL string

func Init(url string) *sql.DB {
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
