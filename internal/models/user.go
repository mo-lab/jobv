package models

import "time"

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Phone     string    `json:"phone" bson:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Role      string    `json:"-" bson:"role,omitempty"`
}
