package dtos

import "time"

type ProfileCardData struct {
	ID           	  string  `json:"id"`
	Username     	  string  `json:"username"`
	DisplayName  	  string  `json:"displayName"`
	AvatarURL    	  *string `json:"avatarUrl,omitempty"`
	BannerURL    	  *string `json:"bannerUrl,omitempty"`
	Rating       	  float64 `json:"rating"`
	ReviewsCount 	  int     `json:"reviewsCount"`
	SessionsPlayed    int     `json:"sessionsPlayed"`
	SessionsHosted    int     `json:"sessionsHosted"`
	PreferredFormat   *SessionFormat `json:"preferredFormat,omitempty"`
	PreferredType     *SessionType   `json:"preferredType,omitempty"`
	IsFollowed     	  bool    `json:"isFollowed"`
}

type SystemStat struct {
	GameSystem
	SessionsCount int `json:"sessionsCount"`
}

type NextSession struct {
	ScheduledAt time.Time     `json:"scheduledAt"`
	Format      SessionFormat `json:"format"`
	Type        SessionType   `json:"type"`
	System      GameSystem `json:"system"`
}

type SessionCardData struct {
	UserID            string          `json:"userId"`
	SystemStats       []SystemStat    `json:"systemStats,omitempty"`
	NextSession       *NextSession    `json:"nextSession,omitempty"`
}
