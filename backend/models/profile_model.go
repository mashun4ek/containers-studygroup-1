package models

import "time"

// Profile model
type Profile struct {
	ID        int       `json:"id"`
	DOB       time.Time `json:"dob"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
}
