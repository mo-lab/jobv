package models

type OTP struct {
	Phone string `json:"phone,omitempty"`
	Token string `json:"code,omitempty"`
}
