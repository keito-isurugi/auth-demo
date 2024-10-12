package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	username = "admin"
	password = "password"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーを取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// 認証がない場合、認証を要求
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// "Basic " プレフィックスを取り除く
		authHeaderParts := strings.SplitN(authHeader, " ", 2)
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Base64デコード
		decoded, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// ユーザー名とパスワードのチェック
		pair := strings.SplitN(string(decoded), ":", 2)
		if len(pair) != 2 || pair[0] != username || pair[1] != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 認証が成功したら次のハンドラに進む
		next(w, r)
	}
}
