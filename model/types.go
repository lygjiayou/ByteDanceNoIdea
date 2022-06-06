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

// lyg
type UserResponse struct {
	Response
	User UserInfo `json:"user"`
}

// qianyu
// 发布列表，以及点赞列表用到这个了（点赞列表之一）
type PublishListResponse struct {
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoInfo `json:"video_list"`  // 用户发布的视频列表
}

// qianyu
type PublishListRequest struct {
	UserID int64 `json:"user_id"`
}
