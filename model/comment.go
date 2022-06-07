package model

type Comment struct {
	ID          int64  `gorm:"column:id;autoIncrement;primaryKey"`           // 视频评论ID
	VideoID     int64 `gorm:"column:video_id"`     // 视频ID
	CommentUserID int64 `gorm:"column:comment_user_id"` // 评论用户ID
	Content     string `gorm:"column:content"`      // 评论内容
	CreateTime  int64 `gorm:"column:create_time"`  // 评论发布日期
}

func (Comment) TableName() string {
	return "comments"
}

type CommentRequestParam struct {
	VideoID int64 `json:"video_id"`
	ActionType	int `json:"action_type"`
	CommentText string `json:"comment_type"`
	CommentID int64	`json:"comment_id"`
}

type CommentInfo struct {
	ID int64 `json:"id"`
	Author UserInfo `json:"user"`
	Content string	`json:"content"`
	CreateTime int64	`json:"create_date"`
}

type CommentActionResponse struct {
	Response
	CommentReturn CommentInfo `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []CommentInfo `json:"comment_list,omitempty"`
}