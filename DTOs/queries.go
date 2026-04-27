package dtos

type UserListSort string

const (
	SortRating     UserListSort = "rating"
	SortRecent     UserListSort = "recent"
	SortFollowedAt UserListSort = "followedAt"
)