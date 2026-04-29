package dtos

type UserListSort string

const (
	SortRating       UserListSort = "rating"
	SortRecent       UserListSort = "recent"
	SortFollowedAt   UserListSort = "followedAt"
	SortReviewsCount UserListSort = "reviews"
)

type SortOrder string

const (
	SortAsc  SortOrder = "ASC"
	SortDesc SortOrder = "DESC"
)