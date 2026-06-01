package dtos

type SessionListResponse struct {
	Items      []Session            `json:"items"`
	NextCursor string               `json:"nextCursor,omitempty"`
	Users      map[string]UserBrief `json:"users"`
}

type SessionResponse struct {
	Session Session              `json:"session"`
	Players []SessionPlayer      `json:"players"`
	Users   map[string]UserBrief `json:"users"`
	CampaignRef *SessionCampaignRef `json:"campaign"`
}

type SessionPlayersResponse struct {
	Players []SessionPlayer      `json:"players"`
	Users   map[string]UserBrief `json:"users"`
}

type SessionCampaignRef struct {
	CampaignID string `json:"campaignId"`
	Title string `json:"title"`
	Index int `json:"index"` 
}