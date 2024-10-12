package main

import (
	"fmt"
	"net/http"
    "github.com/keito-isurugi/security-demo/view"
    "github.com/keito-isurugi/security-demo/middleware"
    "github.com/keito-isurugi/security-demo/handler"
)

func secret(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You have accessed the secret content!")
}

func main() {
    http.HandleFunc("/secret", middleware.BasicAuth(secret))
    http.HandleFunc("/login", view.FormHandler)
    http.HandleFunc("/id_pass_auth", middleware.Post(handler.IdPassAuthHandler))


    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

