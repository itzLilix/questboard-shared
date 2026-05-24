package dtos

type UserBrief struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	DisplayName string  `json:"displayName"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Rating float64 `json:"rating"`
	Hosted int `json:"hosted"`
	Played int `json:"played"`
}
