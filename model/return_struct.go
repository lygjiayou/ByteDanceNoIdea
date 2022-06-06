package model

// VideoInfo publishlist 需要返回给用户的Video信息
type VideoInfo struct {
	ID            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

// UserInfo 展示所需的用户信息
type UserInfo struct {
	ID            int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
