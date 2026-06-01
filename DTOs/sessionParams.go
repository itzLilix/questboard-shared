package dtos

type SessionFormat string

const (
	Online  SessionFormat = "online"
	Offline SessionFormat = "offline"
)

func (f SessionFormat) Valid() bool {
	switch f {
	case Online, Offline:
		return true
	}
	return false
}

func (f *SessionFormat) UnmarshalText(text []byte) error {
	return unmarshalEnum(f, text, SessionFormat.Valid, "session format")
}

type SessionType string

const (
	OneshotType  SessionType = "oneshot"
	CampaignType SessionType = "campaign"
)

func (t SessionType) Valid() bool {
	switch t {
	case OneshotType, CampaignType:
		return true
	}
	return false
}

func (t *SessionType) UnmarshalText(text []byte) error {
	return unmarshalEnum(t, text, SessionType.Valid, "session type")
}

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

func (a SessionAvailability) Valid() bool {
	switch a {
	case Open, Application, Private:
		return true
	}
	return false
}

func (a *SessionAvailability) UnmarshalText(text []byte) error {
	return unmarshalEnum(a, text, SessionAvailability.Valid, "session availability")
}

type SessionStatus string

const (
	Draft     SessionStatus = "draft"
	Published SessionStatus = "published"
	Ongoing   SessionStatus = "ongoing"
	Completed SessionStatus = "completed"
	Cancelled SessionStatus = "cancelled"
)

func (s SessionStatus) Valid() bool {
	switch s {
	case Draft, Published, Ongoing, Completed, Cancelled:
		return true
	}
	return false
}

func (s *SessionStatus) UnmarshalText(text []byte) error {
	return unmarshalEnum(s, text, SessionStatus.Valid, "session status")
}

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

func (s SessionScope) Valid() bool {
	switch s {
	case ScopeCatalog, ScopeMastering, ScopePlaying:
		return true
	}
	return false
}

func (s *SessionScope) UnmarshalText(text []byte) error {
	return unmarshalEnum(s, text, SessionScope.Valid, "session scope")
}

// StatusPresetPublic is a meta-value accepted by the status filter that the
// server expands to {Published, Ongoing, Completed}.
const StatusPresetPublic = "public"