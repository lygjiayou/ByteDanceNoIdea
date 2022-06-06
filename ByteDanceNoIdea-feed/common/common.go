package common

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//Video表
type Video struct {
	Id             int64      `json:"id,omitempty"`
	Author         User       `json:"author" gorm:"-"`
	PlayUrl        string     `json:"play_url,omitempty"`
	CoverUrl       string     `json:"cover_url,omitempty"`
	FavoriteCount  int64      `json:"favorite_count,omitempty"`
	CommentCount   int64      `json:"comment_count,omitempty"`
	IsFavorite     bool       `json:"is_favorite,omitempty"`
	PublisherToken string     `json:"publisher_token"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

//User表
type User struct {
	Id            int64      `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	FollowCount   int64      `json:"follow_count,omitempty" `
	FollowerCount int64      `json:"follower_count,omitempty" `
	IsFollow      bool       `json:"is_follow,omitempty" `
	Password      string     `json:"password,omitempty"`
	Token         string     `json:"token,omitempty" gorm:"unique"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
