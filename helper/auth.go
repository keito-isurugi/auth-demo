package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// パスワードをハッシュ化する関数
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPasswordはハッシュ化されたパスワードを返す
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// ハッシュを文字列に変換して返す
	return string(hash), nil
}

// パスワードとハッシュの比較を行う関数
func CheckPasswordHash(password, hash string) bool {
	// bcrypt.CompareHashAndPasswordは、パスワードが一致するかを確認
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


var secretKey = []byte("mysecret")

// トークン生成関数
func GenerateJWT(userID string) (string, error) {
	fmt.Println("userID", userID)
	// トークンのペイロード（クレーム）
	claims := jwt.MapClaims{
		"sub": userID,                           // ユーザーID
		"exp": time.Now().Add(time.Minute).Unix(), // 有効期限
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークン署名
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// トークンのパースと検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error)  {
		// 署名方式が正しいか確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// エラーがある場合、期限切れエラーかどうかを確認
	if err != nil {
		// エラーメッセージが期限切れの場合
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			return token, fmt.Errorf("token is expired")
		}
		// それ以外のエラーはそのまま返す
		return nil, err
	}

	// トークンが有効かチェック
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
