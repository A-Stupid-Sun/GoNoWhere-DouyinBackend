package dao

import (
	"douyin/config"
	"douyin/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// 初始化数据库、数据表的迁移、设置连接池数量
// 创建对应数据的钩子操作，比如创建用户时，在保存密码之前执行加密操作

// InitDB 显式初始化，连接数据库，表迁移等
var db *gorm.DB

func InitDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassWord,
		config.DbHost,
		config.DbPort,
		config.DbName)

	fmt.Println(dns)
	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //关闭外键！！！
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,       //默认在表的后面加s
			TablePrefix:   "t_douyin_", // 表名前缀
		},
		SkipDefaultTransaction: true, // 禁用默认事务
	})

	if err != nil {
		log.Println("数据库连接失败,err:", err)
	}

	err = db.AutoMigrate(&model.Video{}, &model.User{}, &model.UserLogin{}, &model.Follow{}, &model.Comments{}, &model.Favorite{}) //TODO 数据库自动迁移
	if err != nil {
		log.Println("数据库自动迁移失败，err:", err)
	}
	sqlDb, _ := db.DB()

	// TODO 这方面后期再说吧，参数到底设为多少
	sqlDb.SetMaxIdleConns(50)                  // 连接池中的最大闲置连接数
	sqlDb.SetMaxOpenConns(100)                 // 数据库的最大连接数量
	sqlDb.SetConnMaxLifetime(10 * time.Second) // 连接的最大可复用时间
}
