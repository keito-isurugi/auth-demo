package helper

import "golang.org/x/crypto/bcrypt"

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