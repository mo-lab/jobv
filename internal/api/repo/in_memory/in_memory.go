package repo

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Entry represents a key-value pair with a TTL
type Entry struct {
	Phone      string
	Value      string
	Expiration time.Time
}
type ReqLog struct {
	VerifyCounter int       
	ResetTime     time.Time 
}


// In-memory Storage
var Storage = make(map[string]Entry)
var ReqLogger = make(map[string]ReqLog)
var Mu sync.Mutex // Mutex to handle concurrent access
func NewEntary(phone, value string) Entry {

	expiration := time.Now().Add(2 * time.Minute)
	return Entry{
		Phone:      phone,
		Value:      value,
		Expiration: expiration,
		
	}
}

// Function to store data with TTL
func StoreData(key, value string, ttl time.Duration) {
	Mu.Lock() // Lock the mutex to prevent concurrent writes
	defer Mu.Unlock()
	Storage[key] = Entry{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
	fmt.Printf("verify code for number: %s is code: %s \n", key, value)
}

// Function to retrieve data
func RetrieveData(key string) (string, error) {
	Mu.Lock() // Lock the mutex to prevent concurrent reads
	defer Mu.Unlock()
	entry, exists := Storage[key]
	if !exists {
		return "", errors.New("Key not found")
	}
	if time.Now().After(entry.Expiration) {
		delete(Storage, key) // Remove expired entry
		return "", errors.New("Key expaired")
	}
	return entry.Value, nil
}

// Function to clean up expired entries
func Cleanup() {
	Mu.Lock()
	defer Mu.Unlock()
	for key, entry := range Storage {
		if time.Now().After(entry.Expiration) {
			delete(Storage, key)
		}
	}
}
func RunCleanUp(d time.Duration) {
	if d.Seconds() <= 0 {
		d = 2 * time.Minute
	}
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			Cleanup()
		}
	}()
}