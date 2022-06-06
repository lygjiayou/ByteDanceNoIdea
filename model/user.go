package model

import (
	"ByteDanceNoIdea/utils/errmsg"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

// 基础接口-用户信息
type User struct {
	ID            int64  `gorm:"type:int; not null" json:"id"`                        // 用户id
	UserName      string `gorm:"type:varchar(20); not null" json:"username"`          // 用户名称
	Password      string `gorm:"type:varchar(20); not null" json:"password"`          // 用户密码
	FollowCount   int64  `gorm:"type:int; not null default:0 " json:"follow_count"`   // 关注总数
	FollowerCount int64  `gorm:"type:int; not null default:0 " json:"follower_count"` // 粉丝总数
	//Gender        string `gorm:"type:varchar(6); not null" json:"gender"`    // male-男性，female-女性
	//IsFollow      bool   `gorm:"column:is_follow"`     // true-已关注，false-未关注
}

// 返回UserInfo的信息和发布列表的信息
type UserInfo struct {
	ID            int64  `json:"id"`           // 用户id
	UserName      string `json:"username"`     // 用户名称		// 用户密码
	FollowCount   int64  `json:"follow_count"` // 关注总数
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User

	Db.Select("id").Where("user_name = ?", name).First(&user)

	if user.ID > 0 {
		fmt.Println("CheckUser Error")
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) (id int64, code int) {
	db := Db.Create(&data)
	err := db.Error
	if err != nil {
		fmt.Println("createError")
		return -1, errmsg.ERROR
	}
	return data.ID, errmsg.SUCCESS
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User

	Db.Where("user_name = ?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}

	return errmsg.SUCCESS
}

// 密码加密
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPW(u.Password)
	return
}

func ScryptPW(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 根据username获取到user信息
func FindUserInfo(user User) (int, error) {
	find := Db.Select("id", "user_name", "follow_count", "follower_count").Where("username=?", user.UserName).Find(user) // 已经写入到user里了
	return int(find.RowsAffected), find.Error
}

// 根据username查询userID
func (user *User) FindByUsername() (int64, error) {
	//var user User
	result := Db.Where("user_name=?", user.UserName).Select("id").Find(&user)
	return user.ID, result.Error
}

// FindByUserID 通过用户ID查找用户
func (user *User) FindByUserID(userid string) (int, error) {
	find := Db.Select("id", "user_name", "follow_count", "follower_count").Where("id=?", userid).Find(user)
	return int(find.RowsAffected), find.Error
}

// qianyu FindUserInfoByID 查找author所有信息
func (user *User) FindUserInfoByID() error {
	find := Db.Select("id", "user_name", "follow_count", "follower_count").Where("id = ?", user.ID).Find(user)
	return find.Error
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
