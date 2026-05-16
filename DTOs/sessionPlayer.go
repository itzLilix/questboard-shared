package dtos

import "time"

type SessionPlayerStatus string

const (
	PlayerActive SessionPlayerStatus = "active"
	PlayerKicked SessionPlayerStatus = "kicked"
	PlayerLeft   SessionPlayerStatus = "left"
)

type SessionPlayer struct {
	SessionID string              `json:"sessionId"`
	PlayerID  string              `json:"playerId"`
	Status    SessionPlayerStatus `json:"status"`
	JoinedAt  time.Time           `json:"joinedAt"`
	Character *Character          `json:"character,omitempty"`
}
