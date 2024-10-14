package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/keito-isurugi/auth-demo/helper"
	"github.com/keito-isurugi/auth-demo/infra"
	"gorm.io/gorm"
)

func IdPassAuthHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.FormValue("id"))
		password, _ := helper.HashPassword(r.FormValue("password"))
		
		user, err := infra.GetUser(db, id)
		if err != nil {
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
			return
		}
		fmt.Println("===========================")
		fmt.Println(password)
		fmt.Println(user.Password)

		if helper.CheckPasswordHash(user.Password, password) {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		// 認証成功時にセッションIDを生成してクッキーに保存
		sessionID := uuid.New()
		expiration := time.Now().Add(1 * time.Hour)
		
		if err := infra.SaveSession(db, id, sessionID, expiration); err != nil {
			http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
			return
		}

		// クッキーにセッションIDを保存
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   sessionID.String(),
			Expires: expiration,
		}
		http.SetCookie(w, cookie)

		w.Write([]byte("Authentication succeeded"))
	}
}
