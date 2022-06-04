package model

type Video struct {
	Id            int64  `gorm:"column:id"`             // 视频ID
	Title         string `gorm:"column:title"`          // 视频标题
	AuthorID      string `gorm:"column:author_id"`      // 视频作者ID
	PlayUrl       string `gorm:"column:play_url"`       // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url"`      // 	视频封面地址
	FavoriteCount string `gorm:"column:favorite_count"` // 视频点赞总数
	CommentCount  string `gorm:"column:comment_count"`  // 视频评论总数
	IssueTime     string `gorm:"column:issue_time"`     // 投稿时间
	IsFavorite    bool   `gorm:"column:is_favorite"`    //是否点赞
}

//type Video struct {
//	Id            int64  `json:"id,omitempty"`
//	Author        User   `json:"author"`
//	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
//	CoverUrl      string `json:"cover_url,omitempty"`
//	FavoriteCount int64  `json:"favorite_count,omitempty"`
//	CommentCount  int64  `json:"comment_count,omitempty"`
//	IsFavorite    bool   `json:"is_favorite,omitempty"`
//}

func (Video) TableName() string {
	return "videos"
}

// GetPublishList 获取发布列表
func GetPublishList(userId int64) []Video {
	var videos []Video
	db.Table("videos").Where("author_id = ?", userId).Find(&videos)
	return videos
}
