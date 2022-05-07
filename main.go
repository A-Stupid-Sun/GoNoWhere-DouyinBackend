package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Boy struct {
	gorm.Model `gorm:"index"`
	Name       string
	ID         int
	Friends    []Friend `gorm:"foreignKey:BoyID;references:ID"`
}

type Friend struct {
	Name  string
	BoyID int
}

func InitDb() {
	dns := "test01:Douyintest01.@tcp(47.96.30.253:3306)/test01?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "t_",
		},
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Println("数据库连接失败,err:", err)
	}

	err = db.AutoMigrate(&Boy{}, &Friend{}) //TODO 数据库自动迁移
	//err = db.AutoMigrate(&Dog{}, &Girl{}, &Info{})
	//err = db.AutoMigrate(&Custom{})
	if err != nil {
		log.Println("数据库自动迁移失败，err:", err)
	}
	sqlDb, _ := db.DB()

	sqlDb.SetMaxIdleConns(10)                  // 连接池中的最大闲置连接数
	sqlDb.SetMaxOpenConns(100)                 // 数据库的最大连接数量
	sqlDb.SetConnMaxLifetime(10 * time.Second) // 连接的最大可复用时间
}

func main() {
	fmt.Println("hello world")
	InitDb()
}
