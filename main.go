package main

import (
	"SmallMall/models"
	"SmallMall/routes"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in JWT授权(数据将在请求头中进行传输) 直接在下框中输入Bearer {token}（注意两者之间是一个空格）\"
// @name Authorization
func main() {
	// 引用数据库
	models.InitDB()
	// 引入路由组件
	routes.RouterInit()
}
