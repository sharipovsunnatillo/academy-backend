package main

import (
	"github.com/sharipov/sunnatillo/academy-backend/pkg/middleware"
	"log"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("GET /api/v1/admin"))
	if err != nil {
		log.Println("Warn: ", err)
	}
}

func POST(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("POST /api/v1/admin"))
	if err != nil {
		log.Println("Warn: ", err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/admin", GET)
	mux.HandleFunc("POST /api/v1/admin", POST)

	err := http.ListenAndServe(":8080", middleware.Logging(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
