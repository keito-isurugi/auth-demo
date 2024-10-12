package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

const (
	ID       = "id"
	PassWord = "1a1dc91c907325c69271ddf0c944bc72" // "pass"をmd5でハッシュ化した値
)

func IdPassAuthHandler(w http.ResponseWriter, r *http.Request) {
	// フォームデータの取得（例：id, password）
	id := r.FormValue("id")
	password := r.FormValue("password")

	passHash := md5.Sum([]byte(password))

	if ID != id || PassWord != fmt.Sprintf("%x", passHash) {
		w.Write([]byte("Error Authentication failed"))
	}

	w.Write([]byte("Authentication succeeded"))
}
