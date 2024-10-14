package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID    int       `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
