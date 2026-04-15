package dtos

type PublicProfile struct {
	ID             string  `json:"id"`
	Username       string  `json:"username"`
	DisplayName    string  `json:"displayName"`
	AvatarURL      *string `json:"avatarUrl,omitempty"`
	BannerURL      *string `json:"bannerUrl,omitempty"`
	SessionsPlayed int     `json:"sessionsPlayed"`
	SessionsHosted int     `json:"sessionsHosted"`
	Rating         float64 `json:"rating"`
	ReviewsCount   int     `json:"reviewsCount"`
	Bio            *string `json:"bio,omitempty"`
	Links          []Link  `json:"links,omitempty"`
}