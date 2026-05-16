package dtos

import "time"

type CampaignStatus string

const (
	CampaignActive    CampaignStatus = "active"
	CampaignCompleted CampaignStatus = "completed"
	CampaignCancelled CampaignStatus = "cancelled"
	CampaignPaused    CampaignStatus = "paused"
)

type Campaign struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	MasterID    string         `json:"masterId"`
	System      GameSystem     `json:"system"`
	Status      CampaignStatus `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type CampaignSessionTie struct {
	CampaignID        string    `json:"campaignId"`
	SessionID         string    `json:"sessionId"`
	OrderIndex        int       `json:"orderIndex"`
	CachedTitle       string    `json:"cachedTitle"`
	CachedScheduledAt time.Time `json:"cachedScheduledAt"`
	BriefDescription  *string   `json:"briefDescription,omitempty"`
}
