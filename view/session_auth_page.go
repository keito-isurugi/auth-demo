package view

import (
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/keito-isurugi/auth-demo/infra"
	"gorm.io/gorm"
)

func SessionAuthPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Not authenticated valid cookie", http.StatusUnauthorized)
			return
		}

		sessionID, _ := uuid.Parse(cookie.Value)
		_, valid, err := ValidateSession(db, sessionID)
		if err != nil || !valid {
			http.Error(w, "再度ログインしてください", http.StatusUnauthorized)
			return
		}

		tpl := template.HTML(`
			<h1>ログインが必要なページ アクセス成功！</h1>
		`)
		t := template.Must(template.New("a").Parse(`<div>{{ . }}</div>`))
		t.Execute(w, tpl)
	}
}

// セッションを検証する関数
func ValidateSession(db *gorm.DB, sessionID uuid.UUID) (int, bool, error) {
	// セッションIDを取得
	sessions, err := infra.GetSession(db, sessionID)
	if err != nil {
		return 0, false, err
	}

	// 有効期限をチェック
	now := time.Now()
	if now.After(sessions.ExpiresAt) {
		// 有効期限が過ぎている場合はセッションIDを削除する
		_ = infra.DeleteSession(db, sessionID)
		return 0, false, nil
	}

	return sessions.UserID, true, nil
}
