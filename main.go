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

	// ユーザー一覧
	http.HandleFunc("/users", handler.ListUsers(db.DB))

	// Basic認証
	http.HandleFunc("/secret", middleware.BasicAuth(secret))

	/*
	* セッション方式
	*/
	// ログインページ
	http.HandleFunc("/view/login", view.SessionLoginPage)
	// 認可が必要なページ
	http.HandleFunc("/view/session_auth_page", view.SessionAuthPage(db.DB))
	// パスワードリセットをリクエストするページ
	http.HandleFunc("/view/request_password_reset", view.ViewRequestPasswordResetHandler)
	// パスワードリセットページ
	http.HandleFunc("/view/password_reset", view.ViewPasswordResetHandler)

	// ログイン処理
	http.HandleFunc("/session_login", middleware.Post(handler.SessionLoginHandler(db.DB)))
	// ログアウト処理
	http.HandleFunc("/session_logout", handler.SessionLogoutHnadler(db.DB))
	// パスワードリセットメール送信処理
	http.HandleFunc("/request_password_reset", middleware.Post(handler.RequestPasswordReset(db.DB)))
	// パスワードリセット処理
	http.HandleFunc("/password_reset", middleware.Post(handler.PasswordReset(db.DB)))


	/*
	* トークン方式
	*/
	// jwt生成
	http.HandleFunc("/get_jwt", handler.GetToken)
	http.HandleFunc("/valid_jwt", handler.ValidToken)

	// ログインページ
	http.HandleFunc("/view/jwt_login", view.JWTLoginPage)
	// 認可が必要なページ
	http.HandleFunc("/view/jwt_auth_page", middleware.JWTAuth(view.JWTAuthPage(db.DB), db.DB))
	
	// ログイン処理
	http.HandleFunc("/jwt_login", middleware.Post(handler.JWTLogin(db.DB)))


	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
