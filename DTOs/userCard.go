package dtos

import "time"

type SessionFormat string
type SessionType string

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
	IsFollowed     	  bool    `json:"isFollowed,omitempty"`
}

type GameSystemRef struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type SystemStat struct {
	GameSystemRef
	SessionsCount int `json:"sessionsCount"`
}

type NextSession struct {
	ScheduledAt time.Time     `json:"scheduledAt"`
	Format      SessionFormat `json:"format"`
	Type        SessionType   `json:"type"`
	System      GameSystemRef `json:"system"`
}

type SessionCardData struct {
	UserID            string          `json:"userId"`
	SystemStats       []SystemStat    `json:"systemStats,omitempty"`
	PreferredFormats  []SessionFormat `json:"preferredFormats,omitempty"`
	PreferredTypes    []SessionType   `json:"preferredTypes,omitempty"`
	NextSession       *NextSession    `json:"nextSession,omitempty"`
}
