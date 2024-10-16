package handler

import (
	"fmt"
	"net/http"

	"github.com/keito-isurugi/auth-demo/helper"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	userID := "1234"
	token, err := helper.GenerateJWT(userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "GenerateJWT error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("JWT:" + token))
}

func ValidToken(w http.ResponseWriter, r *http.Request) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkwODQzNTIsInN1YiI6IjEyMzQifQ._T_IcK0dlgAs1adMrBTb7PyIFzPmcVSCG-wdYwKUyXA"
	token ,err := helper.ValidateJWT(jwt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid jwt error", http.StatusInternalServerError)
		return
	}

	fmt.Println(token)
	w.Write([]byte("valid token"))
}
