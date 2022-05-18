package main

import (
	"douyin/config"
	"douyin/dao"
	"douyin/router"
)

func main() {
	dao.InitDB() //数据库初始化
	r := router.InitRouter()
	r.Run(config.Port)

}
