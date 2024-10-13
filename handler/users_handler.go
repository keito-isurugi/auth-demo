package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

type PasswordResetToken struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // 自動インクリメントのプライマリキー
	UserID    uint           `gorm:"not null" json:"user_id"`            // ユーザーID（usersテーブルへの外部キー）
	Token     string         `gorm:"size:255;not null" json:"token"`     // リセット用トークン
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`         // トークンの有効期限
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`   // トークンの生成日時
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`  // 論理削除用のカラム
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

func RequestPasswordReset(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// トークン生成、有効期限生成
		token := "hoge-token"
		// user_idを受け取る
		// DBにuser_id, トークン、有効期限を保存
		// メール送信
		if err := sendResetEmail(token); err != nil{
			w.Write([]byte("メールの送信に失敗しました。"))
			return
		}
		w.Write([]byte("パスワードリセット用のメールを送信しました。"))
	}
}

func sendResetEmail(token string) error {
    from := "form@example.com"
    to := "to@example.com"
    // password := "your-email-password" // 開発では使わない
    smtpHost := "mailhog"
    smtpPort := "1025"

    subject := "パスワードリセット"
    body := fmt.Sprintf("以下のリンクをクリックしてパスワードをリセットしてください:\nhttp://localhost:8080/view/login?token=%s", token)

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    // 認証なしでメールを送信
    err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, []byte(msg))
    if err != nil {
        fmt.Println("メール送信エラー:", err)
    }

    return err
}