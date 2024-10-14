package infra

import (
	"time"

	"github.com/google/uuid"
	"github.com/keito-isurugi/auth-demo/model"
	"gorm.io/gorm"
)

func SaveSession(db *gorm.DB, id int, sessionID uuid.UUID, expiresAt time.Time) error {
	session := model.Session{
		UserID: id,
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}

	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}

func GetSession(db *gorm.DB, sessionID uuid.UUID) (model.Session, error) {
	var sessions model.Session

	if err := db.Where("session_id", sessionID).First(&sessions).Error; err != nil {
		return model.Session{}, err
	}

	return sessions, nil
}

func DeleteSession(db *gorm.DB, sessionID uuid.UUID) error {
	if err := db.Where("session_id", sessionID).Delete(&model.Session{}).Error; err != nil {
		return err
	}
	return nil
}