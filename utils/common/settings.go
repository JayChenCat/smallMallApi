package common

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Mode     string
	Host     string
	HttpPort string
	JwtKey   string

	DbType     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径后重试！！！")
	}
	ReadServerConfig(file)
	ReadDataBaseConfig(file)
}

//读取Server配置
func ReadServerConfig(file *ini.File) {
	Mode = file.Section("server").Key("Mode").MustString("debug")
	Host = file.Section("server").Key("Host").MustString("127.0.0.1")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("88SWEET&Jay77Love13520")
}

//读取数据库连接配置
func ReadDataBaseConfig(file *ini.File) {
	DbType = file.Section("database").Key("DbType").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
