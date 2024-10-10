package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!!!")
    })
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Bar!")
    })
	fmt.Println("localhost:8080 server runnig ...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}