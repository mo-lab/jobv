package main

import (
	"log"
	"net/http"

	"github.com/mo-lab/jobv/api/v2/internal/api/handlers"
	mw "github.com/mo-lab/jobv/api/v2/internal/api/middlewares"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /otp/send", handlers.SendOTPHandler)
	mux.HandleFunc("POST /otp/verify", handlers.VerifyOTPHandler)
	mux.HandleFunc("GET /users", handlers.GetUsersHandler)
	mux.HandleFunc("GET /users/search", handlers.SearchUsersHandler)
	// Protect the search endpoint with JWT authentication
	mux.HandleFunc("GET /protected/users/search", mw.JwtAuth(handlers.SearchUsersHandler))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
