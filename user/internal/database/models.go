// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Level     int32
	Badges    []string
	IsBanned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
