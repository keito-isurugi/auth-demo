package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keito-isurugi/auth-demo/db"
	"github.com/keito-isurugi/auth-demo/handler"
	"github.com/keito-isurugi/auth-demo/middleware"
	"github.com/keito-isurugi/auth-demo/view"
)

func secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have accessed the secret content!")
}

func main() {
	// データベース接続
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	http.HandleFunc("/secret", middleware.BasicAuth(secret))
	http.HandleFunc("/view/login", view.FormHandler)
	http.HandleFunc("/id_pass_auth", middleware.Post(handler.IdPassAuthHandler))

	http.HandleFunc("/users", handler.ListUsers(db.DB))

	http.HandleFunc("/view/request_password_reset", view.ViewRequestPasswordResetHandler)
	http.HandleFunc("/view/password_reset", view.ViewPasswordResetHandler)
	http.HandleFunc("/request_password_reset", middleware.Post(handler.RequestPasswordReset(db.DB)))
	http.HandleFunc("/password_reset", middleware.Post(handler.PasswordReset(db.DB)))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
