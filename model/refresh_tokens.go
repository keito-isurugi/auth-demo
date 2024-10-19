package model

import (
    "time"
    "github.com/google/uuid"
)

type RefreshToken struct {
    RefreshToken uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID       int      `gorm:"not null"`
    ExpiresAt    time.Time `gorm:"not null"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
}
