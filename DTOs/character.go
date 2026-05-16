package dtos

import "time"

type Character struct {
	ID          string    `json:"id"`
	PlayerID    string    `json:"playerId"`
	Name        string    `json:"name"`
	Class       *string   `json:"class,omitempty"`
	Level       *int      `json:"level,omitempty"`
	AvatarURL   *string   `json:"avatarUrl,omitempty"`
	Description *string   `json:"description,omitempty"`
	SheetURL    *string   `json:"sheetUrl,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
