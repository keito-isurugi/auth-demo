package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func ListUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User

		if err := db.Find(&users).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to encode users", http.StatusInternalServerError)
			return
		}
	}
}
