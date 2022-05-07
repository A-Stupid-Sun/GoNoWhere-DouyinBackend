package config

import "gopkg.in/ini.v1"

// 解析配置文件
var (
	AppMode    string
	Port       string
	JwtKey     string
	Dbtype     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {

}

// LoadServer 加载服务器配置
func LoadServer(file *ini.File) {
	s := file.Section("server")
	AppMode = s.Key("AppMode").MustString("debug")
	Port = s.Key("Port").MustString("3001")
	JwtKey = s.Key("JwtKey").MustString("DouYin")

}

// LoadDb 加载数据库相关配置
func LoadDb(file *ini.File) {
	s := file.Section("database")
	Dbtype = s.Key("Dbtype").MustString("mysql")
	DbName = s.Key("DbName").MustString("test01")
	DbPort = s.Key("DbPort").MustString("DbPort")
	DbHost = s.Key("DbHost").MustString("DbHost")
	DbUser = s.Key("DbUser").MustString("root")
	DbPassWord = s.Key("DbPassWord").MustString("root")
}
