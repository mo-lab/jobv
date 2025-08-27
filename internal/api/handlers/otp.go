package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	repo "github.com/mo-lab/jobv/api/v2/internal/api/repo/in_memory"
	"github.com/mo-lab/jobv/api/v2/internal/models"
)

type ReqLog struct {
	VerifyCounter int
	ResetTime     time.Time
}

func NewLog() ReqLog {

	return ReqLog{
		VerifyCounter: 0,
		ResetTime:     time.Now().Add(10 * time.Minute),
	}
}

var ReqLogger = make(map[string]ReqLog)
var Mu sync.Mutex

func GenerateToken() string {

	return strconv.Itoa(rand.Intn(900000) + 100000)
}
func SendOTPHandler(w http.ResponseWriter, r *http.Request) {
	// parse phone number from request

	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}
	// Generate OTP

	// SMS service to send the OTP
	if IsLimited(phone) {
		http.Error(w, "Too Many requests", http.StatusTooManyRequests)
		return
	}
	otp := GenerateToken()
	fmt.Printf("Sending OTP %s to phone %s\n", otp, phone)
	//pub to nats
	repo.StoreData(phone, otp, 2*time.Minute)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP sent successfully"))

}

func IsLimited(phone string) bool {
	lg, exist := ReqLogger[phone]
	if !exist {
		lg = NewLog()
	}
	Mu.Lock()
	defer Mu.Unlock()
	lg.VerifyCounter++
	if lg.ResetTime.Before(time.Now()) {
		lg.ResetTime = time.Now().Add(10 * time.Minute)
		lg.VerifyCounter = 0
	}
	if lg.VerifyCounter > 3 {
		return true
	}
	return false
}
func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {

	var req models.OTP
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	// parse phone number from request
	if req.Phone == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}
	t, err := repo.RetrieveData(req.Phone)
	// Verify OTP
	if err != nil {
		http.Error(w, "Phone number does not exist", http.StatusNotFound)
		return
	}
	if req.Token != t {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}
	
	LoginHandler(w, req.Phone, "user")
}
