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

func GetToken(w http.ResponseWriter, r *http.Request) {
	userID := "1234"
	token, err := helper.GenerateJWT(userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "GenerateJWT error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("JWT:" + token))
}

func ValidToken(w http.ResponseWriter, r *http.Request) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkyNjIyMTAsInN1YiI6IjEyMzQifQ.XRANbDR5Ow12vKXDC62YPSb9BC1cRd7SBzHWD5ciuAk"
	token, err := helper.ValidateJWT(jwt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid jwt error", http.StatusInternalServerError)
		return
	}

	fmt.Println(token)
	w.Write([]byte("valid token"))
}

func JWTLogin(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.FormValue("id"))
		password, _ := helper.HashPassword(r.FormValue("password"))

		user, err := infra.GetUser(db, userID)
		if err != nil {
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
			return
		}

		if helper.CheckPasswordHash(user.Password, password) {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}
		
		// jwt発行
		token, err := helper.GenerateJWT(strconv.Itoa(user.ID))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "GenerateJWT error", http.StatusInternalServerError)
			return
		}
		// クッキーにjwtを保存
		cookie := &http.Cookie{
			Name:  "jwt_auth_key",
			Value: token,
		}
		http.SetCookie(w, cookie)

		// リフレッシュトークン発行、DBに保存
		rt := uuid.New()
		t := time.Now().Add(time.Hour * 2)
		if err := infra.SaveRefreshTokens(db, userID, rt, t); err != nil {
			fmt.Println(err)
			http.Error(w, "リフレッシュトークン登録エラー", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("JWTログイン成功！"))
	}
}
