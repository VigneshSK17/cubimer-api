// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Scramble struct {
	ID        int64
	UserID    int64
	Time      int32
	Scramble  string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type User struct {
	ID       int64
	Username string
	Password string
}
