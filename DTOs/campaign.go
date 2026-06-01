package dtos

import "time"

type CampaignStatus string

const (
	CampaignActive    CampaignStatus = "active"
	CampaignCompleted CampaignStatus = "completed"
	CampaignCancelled CampaignStatus = "cancelled"
	CampaignPaused    CampaignStatus = "paused"
)

func (s CampaignStatus) Valid() bool {
	switch s {
	case CampaignActive, CampaignCompleted, CampaignCancelled, CampaignPaused:
		return true
	}
	return false
}

func (s *CampaignStatus) UnmarshalText(text []byte) error {
	return unmarshalEnum(s, text, CampaignStatus.Valid, "campaign status")
}

type Campaign struct {
	ID           string               `json:"id"`
	Title        string               `json:"title"`
	Description  *string              `json:"description,omitempty"`
	MasterID     string               `json:"masterId"`
	System       *GameSystem          `json:"system"`
	Status       CampaignStatus       `json:"status"`
	Availability SessionAvailability  `json:"availability"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
	Sessions     []CampaignSessionTie `json:"sessions"`
	SessionCount int                  `json:"sessionCount"`
}

type CampaignSessionTie struct {
	CampaignID        string    `json:"-"`
	SessionID         string    `json:"sessionId"`
	OrderIndex        int       `json:"orderIndex"`
	CachedTitle       string    `json:"cachedTitle"`
	CachedScheduledAt time.Time `json:"cachedScheduledAt"`
	BriefDescription  *string   `json:"briefDescription,omitempty"`
}
