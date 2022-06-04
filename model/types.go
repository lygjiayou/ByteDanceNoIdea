package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//type User struct {
//	ID            int64  `gorm:"type:int; not null" json:"id"`
//	UserName      string `gorm:"type:varchar(20); not null" json:"username"`
//	Password      string `gorm:"type:varchar(20); not null" json:"password"`
//	FollowCount   int64  `gorm:"type:int; not null default:0 " json:"follow_count"`
//	FollowerCount int64  `gorm:"type:int; not null default:0 " json:"follower_count"`
//}

type UserResponse struct {
	Response
	User User `json:"user"`
}
