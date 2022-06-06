package model

//favorite数据库模型
type Favorite struct {
	ID	int64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserID	int64 `gorm:"column:user_id"`
	VideoID int64 `gorm:"column:video_id"`
}
func (Favorite) TableName() string {
	return "favorite"
}
//favorite列表
type FavoriteList struct {
	Response
	VideoInfoList []VideoInfo `json:"video_list"`
}