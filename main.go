package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keito-isurugi/security-demo/db"
	"github.com/keito-isurugi/security-demo/handler"
	"github.com/keito-isurugi/security-demo/middleware"
	"github.com/keito-isurugi/security-demo/view"
	"gorm.io/gorm"
)

func secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have accessed the secret content!")
}

func main() {
	// データベース接続
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// ユーザー一覧を取得
	users, err := GetUsers(db.DB)
	if err != nil {
		log.Fatal("Failed to get users:", err)
	}

	// ユーザー一覧を表示
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	http.HandleFunc("/secret", middleware.BasicAuth(secret))
	http.HandleFunc("/login", view.FormHandler)
	http.HandleFunc("/id_pass_auth", middleware.Post(handler.IdPassAuthHandler))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User

	// DBからデータを取得
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}
