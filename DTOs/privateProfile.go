package dtos

import "time"

type PrivateProfile struct {
	PublicProfile
	Email           string     	`json:"email"`
	CreatedAt       time.Time  	`json:"createdAt"`
	LastLogin       time.Time 	`json:"lastLogin"`
	Role            Role       	`json:"role"`
	IsEmailVerified bool       	`json:"isEmailVerified"`
}