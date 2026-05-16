package dtos

import "time"

type SessionCommentary struct {
	ID        string    `json:"id"`
	SessionID string    `json:"sessionId"`
	AuthorID  string    `json:"authorId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
