package infra

import (
	"github.com/keito-isurugi/auth-demo/model"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id int) (model.User, error) {
	var user model.User

	if err := db.Where("id", id).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
