package dtos

type SessionFormat string

const (
	Online  SessionFormat = "online"
	Offline SessionFormat = "offline"
)

type SessionType string

const (
	OneshotType  SessionType = "oneshot"
	CampaignType SessionType = "campaign"
)

type GameSystem struct {
	Id         string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	IsCurated  bool   `json:"isCurated"`
	BadgeColor string `json:"badgeColor"`
}

type SessionAvailability string

const (
	Open        SessionAvailability = "open"
	Application SessionAvailability = "application"
	Private     SessionAvailability = "private"
)

type SessionStatus string

const (
	Draft     SessionStatus = "draft"
	Published SessionStatus = "published"
	Ongoing   SessionStatus = "ongoing"
	Completed SessionStatus = "completed"
	Cancelled SessionStatus = "cancelled"
)

type Location struct {
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type SessionScope string

const (
	ScopeCatalog   SessionScope = "catalog"
	ScopeMastering SessionScope = "mastering"
	ScopePlaying   SessionScope = "playing"
)

// StatusPresetPublic is a meta-value accepted by the status filter that the
// server expands to {Published, Ongoing, Completed}. Not a real SessionStatus
// — the domain enum stays pure; this constant just documents the wire string.
const StatusPresetPublic = "public"