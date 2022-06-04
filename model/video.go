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
}

//获取发布列表

func GetPublishList(userId string) []Video {
	var videos []Video
	err := db.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {

	}

}
