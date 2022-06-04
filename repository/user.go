package repository

// 基础接口-用户信息
type User struct {
	ID            int64  `gorm:"column:id"`             // 用户id
	UserName      string `gorm:"column:username"`       // 用户名称
	Password      string `gorm:"column:password"`       // 用户密码
	FollowCount   int64  `gorm:"column:follow_count"`   // 关注总数
	FollowerCount int64  `gorm:"column:follower_count"` // 粉丝总数
}

//func (User) TableName() string {
//	return "user"
//}
//
//type UserDao struct {
//}
//
//var userDao *UserDao
//var userOnce sync.Once
//
//func NewUserDaoInstance() *UserDao {
//	userOnce.Do(
//		func() {
//			userDao = &UserDao{}
//		})
//	return userDao
//}
//
//func (*UserDao) QueryUserById(id int64) (*User, error) {
//	var user User
//	err := db.Where("id = ?", id).Find(&user).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil, nil
//	}
//	if err != nil {
//		//util.Logger.Error("find user by id err:" + err.Error())
//		return nil, err
//	}
//	return &user, nil
//}
//
//func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
//	var users []*User
//	err := db.Where("id in (?)", ids).Find(&users).Error
//	if err != nil {
//		//util.Logger.Error("batch find user by id err:" + err.Error())
//		return nil, err
//	}
//	userMap := make(map[int64]*User)
//	for _, user := range users {
//		userMap[user.ID] = user
//	}
//	return userMap, nil
//}
