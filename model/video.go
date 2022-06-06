package model

// yanghai
// Video数据库模型
type Video struct {
	ID            int64  `gorm:"column:id;autoIncrement;primaryKey"`
	AuthorID      int64  `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	CreateTime    int64  `gorm:"column:create_time"`
	Title         string `gorm:"column:title"`
}

func (Video) TableName() string {
	return "videos"
}

// GetPublishList 获取发布列表
func GetPublishList(userId int64) []Video {
	var videos []Video
	Db.Where("author_id = ?", userId).Find(&videos)
	return videos
}
