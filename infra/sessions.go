package infra

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
    UserID    int       `json:"user_id"`
    SessionID uuid.UUID `json:"session_id"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
}

func SaveSession(db *gorm.DB, id int, sessionID uuid.UUID, expiresAt time.Time) error {
	session := Session{
		UserID: id,
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}

	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}
