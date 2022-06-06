package model

//// 基础接口-用户信息
type User struct {
	ID            int64  `gorm:"column:id"`             // 用户id
	UserName      string `gorm:"column:username"`       // 用户名称
	Password      string `gorm:"column:password"`       // 用户密码
	FollowCount   int64  `gorm:"column:follow_count"`   // 关注总数
	FollowerCount int64  `gorm:"column:follower_count"` // 粉丝总数
	//IsFollow      bool   `gorm:"column:is_follow"`     // true-已关注，false-未关注
}

func (User) TableName() string {
	return "user"
}

// userList 存放用户的映射，用以判断用户是否存在。
var userList = make(map[int]bool, 10)

// FindByUsername 第一次登录或者注册时，根据请求的Name查询数据库中的ID，返回
func (user *User) FindByUsername() (int64, error) {
	//var user User
	result := Db.Where("username=?", user.UserName).Select("id").Find(&user)
	return user.ID, result.Error
}

// FindByUserID 第一次登录时通过用户名和密码查找用户，后面访问需要携带token来访问其他资源，token存在服务器（本地）
func (user *User) FindByUserID() (int, error) {
	find := Db.Select("username", "password").Find(user) // name:用户名，password:密码
	return int(find.RowsAffected), find.Error
}

// FindByUsernamePassword 通过用户名和密码查找用户是否存在
func (user *User) FindByUsernamePassword() (int, error) {
	find := Db.Where("username=? AND password=?", user.UserName, user.Password).Find(user)
	//find := db.Where(&User{Name: user.Name, Password: user.Password}).First(&user)
	return int(find.RowsAffected), find.Error
}

// FindByUsername 通过用户名查找用户是否存在
func (user *User) FindByName() (int, error) {
	find := Db.Where("username=?", user.UserName).Find(user)
	return int(find.RowsAffected), find.Error
}

// CreateUser 注册用户
func (user *User) CreateUser() error {
	err := Db.Create(user).Error
	return err
}

// 根据token里解析出的name查询user信息并返回
func (user *User) FindByTokename() (int, error) {
	//var user User
	//result := db.Where("name=?", user.Name).Select("id").Find(&user)
	find := Db.Select("id", "username", "follow_count", "follower_count").Where("username=?", user.UserName).Find(user) // 已经写入到user里了
	return int(find.RowsAffected), find.Error
}

//// 根据token里解析出的name查询user信息并返回
//func (user *User) FindByTokename() (int, error) {
//	//var user User
//	//result := db.Where("name=?", user.Name).Select("id").Find(&user)
//	find := db.Select("id", "username", "follow_count", "follower_count").Where("username=?", user.UserName).Find(user) // 已经写入到user里了
//	return int(find.RowsAffected), find.Error
//}
