package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/keito-isurugi/auth-demo/helper"
	"github.com/keito-isurugi/auth-demo/infra"
	"gorm.io/gorm"
)

func JWTAuth(next http.HandlerFunc, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorization ヘッダーからトークンを取得
		// authHeader := r.Header.Get("Authorization")
		// fmt.Println(authHeader)
		// if authHeader == "" {
		// 	http.Error(w,"Missing Authorization Header", http.StatusUnauthorized)
		// 	return
		// }

		// "Bearer トークン" 形式なので、"Bearer " を取り除く
		// tokenString := authHeader[len("Bearer "):]
		cookie, err := r.Cookie("jwt_auth_key")
		if err != nil {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		fmt.Println(tokenString)

		// トークンの検証
		token, err := helper.ValidateJWT(tokenString)

		if err != nil {
			// リフレッシュトークンの検証
			// リフレッシュトークンが存在していて、有効期限が切れていなければjwtを再発行
			userID, err := GetUserID(token)
			if err != nil {
				http.Error(w, "jwtからユーザーIDが取得できません", http.StatusUnauthorized)
				return
			}

			id, _ := strconv.Atoi(userID)
			rt, err := infra.GetRefreshToken(db, id)
			if err != nil {
				http.Error(w, "リフレッシュトークンが取得できません", http.StatusUnauthorized)
				return
			}

			// リフレッシュトークンの有効期限が切れている場合はログインを促す
			now := time.Now()
			if rt.ExpiresAt.Before(now) {
				http.Error(w, "リフレッシュトークンの有効期限切れです。再度ログインが必要です。", http.StatusUnauthorized)
				return
			}

			// リフレッシュトークンの有効期限が切れていない場合はjwtを再発行
			token, err := helper.GenerateJWT(userID)
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

			w.Write([]byte("jwtを更新しました"))
			next(w, r)
		}

		// クレームからユーザーID(sub)を取得
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["sub"].(string)
			fmt.Println(userID)

			// レスポンスにユーザーIDを返す
			w.Write([]byte(fmt.Sprintf("UserID: %s", userID)))
			next(w, r)
		}
	}
}

func GetUserID(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("jwtからユーザーIDが取得できません")
	}
	userID := claims["sub"].(string)
	return userID, nil
}
