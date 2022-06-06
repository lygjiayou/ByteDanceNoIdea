package repository

type Like struct {
	Id      int64  `gorm:"column:id"`    // 点赞记录ID
	UserId  string `gorm:"column:title"` // 点赞用户ID
	VideoId string `gorm:"column:name"`  // 被点赞视频ID
}
