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
}

type SessionPlayersResponse struct {
	Players []SessionPlayer      `json:"players"`
	Users   map[string]UserBrief `json:"users"`
}
