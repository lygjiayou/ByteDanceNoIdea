package model

type Relation struct {
	ID	int64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserID int64	`gorm:"column:user_id"`
	FollowerID	int64	`gorm:"column:follower_id"`
}

type RelationUserInfo struct {
	Response
	UserInfos []UserInfo	`json:"user_list"`
}
