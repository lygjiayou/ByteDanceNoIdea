package model

type Follow struct {
	Id         int64  `gorm:"column:id"`    // 关注记录ID
	FollowerId string `gorm:"column:title"` // 关注者用户ID
	StarId     string `gorm:"column:name"`  // 被关注者用户ID
}
