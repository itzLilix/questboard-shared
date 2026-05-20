package dtos

type UserListSort string

const (
	SortRating        UserListSort = "rating"
	SortRecent        UserListSort = "recent"
	SortFollowedAt    UserListSort = "followedAt"
	SortReviewsCount  UserListSort = "reviews"
	SortSessionsCount UserListSort = "sessions"
)

type SortOrder string

const (
	SortAsc  SortOrder = "ASC"
	SortDesc SortOrder = "DESC"
)

type SessionListSort string

const (
	SortSessionScheduledAt SessionListSort = "scheduled_at"
	SortSessionCreatedAt   SessionListSort = "created_at"
	SortSessionPrice       SessionListSort = "price"
	SortSessionTitle       SessionListSort = "title"
	SortSessionSystem      SessionListSort = "system"
)