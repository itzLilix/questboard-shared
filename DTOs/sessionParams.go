package dtos

type SessionFormat string

const (
	Online  SessionFormat = "online"
	Offline SessionFormat = "offline"
)

type SessionType string

const (
	Oneshot  SessionType = "oneshot"
	Campaign SessionType = "campaign"
)

type GameSystem struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	IsCurated  bool   `json:"isCurated"`
	BadgeColor string `json:"badgeColor"`
}