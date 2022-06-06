package model

import (
	"ByteDanceNoIdea/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//以下是配置mysql数据库
//var (
//	addrMYSQL = "127.0.0.1:3306" //mysql地址
//	account   = "root"           //mysql账号
//	password  = "root"           //mysql密码
//	dbName    = "noideadouyin"   //mysql数据库
//)

var Db *gorm.DB
var err error

// InitMysql 初始化mysql链接
func InitMysql() {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	connString := utils.DbUser + ":" + utils.DbPassWord + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
		//SkipDefaultTransaction: false, //自动开启事务的开关
	})

	// 设置用户信息自动迁移
	err = Db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("自动迁移失败")
	}

	// 设置视频信息自动迁移
	err = Db.AutoMigrate(&Video{})
	if err != nil {
		fmt.Println("自动迁移失败")
	}

	// 设置点赞信息自动迁移
	err = Db.AutoMigrate(&Favorite{})
	if err != nil {
		fmt.Println("自动迁移失败")
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalln("mysql lost:", err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10) // 空闲连接池最大连接数。
	//打开
	sqlDB.SetMaxOpenConns(30)
}
