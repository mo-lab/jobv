package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mo-lab/jobv/api/v2/internal/api/repo/mongodb"
	"github.com/mo-lab/jobv/api/v2/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtSecretKey = []byte("your_secret_key_here")

// /login - POST {phone, role} returns JWT
func LoginHandler(w http.ResponseWriter, phone string, role string) {
	exists := mongodb.CheckUserPhone(phone)
	if !exists {
		mongodb.CreateUser(models.User{Phone: phone, Role: role})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting user details
	q := r.URL.Query()
	id := q.Get("id")
	user, err := mongodb.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	id, err := mongodb.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	ids, err := mongodb.CreateUsers(users)
	if err != nil {
		http.Error(w, "Failed to create users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string][]string{"ids": ids})
}
func SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	phone := q.Get("phone")
	role := q.Get("role")
	filter := bson.M{}

	if phone != "" {
		filter["phone"] = phone
	}
	if role != "" {
		filter["role"] = role
	}
	pg, limit, shouldReturn := CheckPaginationParams(q, w)
	if shouldReturn {
		return
	}
	users, err := mongodb.GetUsers(context.TODO(), filter, pg, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	pg, limit, shouldReturn := CheckPaginationParams(q, w)
	if shouldReturn {
		return
	}
	ctx := context.Background()
	users, err := mongodb.GetUsers(ctx, bson.M{}, pg, limit)
	if err != nil {
		http.Error(w, "iternal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
