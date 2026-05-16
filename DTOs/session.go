package dtos

import "time"

type Session struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Format      SessionFormat `json:"format"`
	ScheduledAt time.Time `json:"scheduledAt"`
	Duration	*float32 `json:"duration"`
	Location    *Location `json:"location"`
	System GameSystem `json:"system"`
	Type SessionType `json:"type"`
	Availability SessionAvailability `json:"availability"`
	Description string `json:"description"`
	PreviewUrl string `json:"previewUrl"`
	MaxSeats uint8 `json:"maxSeats"`
	MasterID string `json:"masterID"`
	Price float32 `json:"price"`
	MasterNotes string `json:"masterNotes"`
	Status SessionStatus `json:"status"`
	FreeSeats uint8 `json:"freeSeats"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}