package dtos

import "time"

type PrivateProfileData struct {
	PublicProfileData
	Email           string     	`json:"email"`
	CreatedAt       time.Time  	`json:"createdAt"`
	LastLogin       time.Time 	`json:"lastLogin"`
	Role            Role       	`json:"role"`
	IsEmailVerified bool       	`json:"isEmailVerified"`
}