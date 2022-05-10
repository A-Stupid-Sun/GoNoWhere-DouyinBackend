package main

import (
	"douyin/config"
	"douyin/dao"
	"douyin/router"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	dao.InitDB() //数据库初始化
	r := router.InitRouter()
	r.Run(config.Port)
}
