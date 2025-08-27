package main

import (
	"log"
	"net/http"

	"github.com/mo-lab/jobv/api/v2/internal/api/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /otp/send", handlers.SendOTPHandler)
	mux.HandleFunc("POST /otp/verify", handlers.VerifyOTPHandler)
	mux.HandleFunc("GET /users", handlers.GetUsersHandler)
	mux.HandleFunc("GET /users/search", handlers.SearchUsersHandler)

	//http.HandleFunc("/users/search", jwtAuth(searchHandler))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
