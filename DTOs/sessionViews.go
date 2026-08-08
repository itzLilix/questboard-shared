package dtos

// SessionMembership is the viewer's relation to a session, shown on catalog
// cards. Sessions the viewer has no relation to are absent from the map.
type SessionMembership string

const (
	MembershipMaster    SessionMembership = "master"
	MembershipPlayer    SessionMembership = "player" // active seat
	MembershipKicked    SessionMembership = "kicked"
	MembershipLeft      SessionMembership = "left"
	MembershipApplicant SessionMembership = "applicant" // pending application
)

type SessionListResponse struct {
	Items      []Session                     `json:"items"`
	NextCursor string                        `json:"nextCursor,omitempty"`
	Users      map[string]UserBrief          `json:"users"`
	Campaigns  map[string]SessionCampaignRef `json:"campaigns"`
	Membership map[string]SessionMembership  `json:"membership"`
}

type SessionResponse struct {
	Session Session              `json:"session"`
	Players []SessionPlayer      `json:"players"`
	Users   map[string]UserBrief `json:"users"`
	CampaignRef *SessionCampaignRef `json:"campaign"`
	Membership SessionMembership  `json:"membership"`
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