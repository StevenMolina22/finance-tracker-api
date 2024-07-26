// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type Author struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

type Transaction struct {
	ID          int64           `json:"id"`
	UserID      sql.NullInt64   `json:"user_id"`
	Amount      sql.NullFloat64 `json:"amount"`
	Date        sql.NullTime    `json:"date"`
	Description sql.NullString  `json:"description"`
	Type        sql.NullString  `json:"type"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	UpdatedAt   sql.NullTime    `json:"updated_at"`
}

type User struct {
	ID        int64        `json:"id"`
	ClerkID   string       `json:"clerk_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
