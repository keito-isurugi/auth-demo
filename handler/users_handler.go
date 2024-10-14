package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"github.com/keito-isurugi/auth-demo/helper"
	"github.com/keito-isurugi/auth-demo/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Token     uuid.UUID `gorm:"not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func ListUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

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
		// user_idに紐づくユーザー取得(送信用メールアドレスに必要)
		var user model.User
		userID := r.FormValue("user_id")
		if err := db.Where("id", userID).First(&user).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
			return
		}

		// トークン生成、有効期限生成
		token := uuid.New()
		expiresAt := time.Now().Add(time.Hour * 1)

		// DBにuser_id, トークン、有効期限をアップサート
		prt := &PasswordResetToken{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: expiresAt,
		}
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"token", "expires_at"}),
		}).Create(&prt).Error; err != nil {
			http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		}

		// メール送信
		if err := sendResetEmail(user.Email, token); err != nil {
			w.Write([]byte("メールの送信に失敗しました。"))
			return
		}
		w.Write([]byte("パスワードリセット用のメールを送信しました。"))
	}
}

func sendResetEmail(toEmail string, token uuid.UUID) error {
	from := "form@example.com"
	to := toEmail
	smtpHost := "mailhog"
	smtpPort := "1025"

	subject := "パスワードリセット"
	body := fmt.Sprintf("以下のリンクをクリックしてパスワードをリセットしてください:\nhttp://localhost:8080/view/password_reset?token=%s", token)

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

func PasswordReset(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// userIDを元にトークン取得
		var prt PasswordResetToken
		userID := r.FormValue("user_id")
		if err := db.Where("id", userID).First(&prt).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
			return
		}

		// トークン、有効期限と比較
		token := r.FormValue("token")
		currentDate := time.Now()
		if prt.Token.String() != token || currentDate.After(prt.ExpiresAt) {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}

		// パスワードをリセットする
		hash, _ := helper.HashPassword(r.FormValue("new_password"))
		user := &model.User{
			ID:       prt.ID,
			Password: hash,
		}
		fmt.Println("===========================")
		fmt.Println(hash)
		fmt.Println(user.Password)
		if err := db.Updates(user).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
		}

		w.Write([]byte("パスワードをリセットしました。"))
	}
}
