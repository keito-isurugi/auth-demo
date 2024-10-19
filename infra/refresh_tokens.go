package infra

import (
	"time"

	"github.com/google/uuid"
	"github.com/keito-isurugi/auth-demo/model"
	"gorm.io/gorm"
)

func SaveRefreshTokens(db *gorm.DB, id int, token uuid.UUID, expiresAt time.Time) error {
	rt := model.RefreshToken{
		RefreshToken: token,
		UserID: id,
		ExpiresAt: expiresAt,
	}

	if err := db.Create(&rt).Error; err != nil {
		return err
	}

	return nil
}

func GetRefreshToken(db *gorm.DB, token uuid.UUID) (model.RefreshToken, error) {
	var rt model.RefreshToken

	if err := db.Where("refresh_token", token).First(&rt).Error; err != nil {
		return model.RefreshToken{}, err
	}

	return rt, nil
}

func DeleteRefreshToken(db *gorm.DB, token uuid.UUID) error {
	if err := db.Where("refresh_token", token).Delete(&model.RefreshToken{}).Error; err != nil {
		return err
	}
	return nil
}