// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID            uuid.UUID      `json:"id"`
	UserID        uuid.UUID      `json:"user_id"`
	GroupID       uuid.NullUUID  `json:"group_id"`
	TypeID        uuid.UUID      `json:"type_id"`
	Name          string         `json:"name"`
	UseDay        time.Time      `json:"use_day"`
	Amount        float64        `json:"amount"`
	Cost          float64        `json:"cost"`
	CostIndicator sql.NullString `json:"cost_indicator"`
	Place         string         `json:"place"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Group struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Type struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGroup struct {
	ID        uuid.UUID     `json:"id"`
	UserID    uuid.NullUUID `json:"user_id"`
	GroupID   uuid.NullUUID `json:"group_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}