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
}

// yanghai
// publishlist返回给用户的Video信息
type VideoInfo struct {
	ID            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

func (Video) TableName() string {
	return "videos"
}

// GetPublishList 获取发布列表
func GetPublishList(userId int64) []Video {
	var videos []Video
	Db.Table("videos").Where("author_id = ?", userId).Find(&videos)
	return videos
}
