package handler

import "net/http"

const (
	ID       = "id"
	PassWord  = "pass"
)

func IdPassAuthHandler(w http.ResponseWriter, r *http.Request) {
	// フォームデータの取得（例：id, password）
    id := r.FormValue("id")
    password := r.FormValue("password")

	if ID != id || PassWord != password {
		w.Write([]byte("Error Authentication failed"))	
	}

	w.Write([]byte("Authentication succeeded"))
}
