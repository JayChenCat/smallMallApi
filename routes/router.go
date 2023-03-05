package routes

import (
	v1 "SmallMall/api/v1"
	docs "SmallMall/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//上面的SmallMall为本项目名称，docs就是swag init自动生成的目录，用于存放 docs.go、swagger.json、swagger.yaml 三个文件。
	"SmallMall/utils/common"
	"SmallMall/utils/middleware"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createInitRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func RouterInit() {
	gin.SetMode(common.Mode)
	router := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，减少性能消耗，上线可设置
	_ = router.SetTrustedProxies(nil)
	//设置head，跨域，日志管理等
	//router.HTMLRender = createInitRender()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	//设置静态文件
	//router.Static("/static", "./web/front/dist/static")
	//router.Static("/admin", "./web/admin/dist")
	//router.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")
	//设置访问前后台的路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	// 浏览器输入可以打开页面
	// http://127.0.0.1:8889/swagger/index.html
	//url := ginSwagger.URL("http://127.0.0.1:3000/swagger/doc.json")
	docs.SwaggerInfo.Title = "美物小程序后台管理系统 API"
	docs.SwaggerInfo.Description = "美物小程序后台管理系统接口"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = common.Host + common.HttpPort //"127.0.0.1:80"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("api/v1")
	{
		/*
		  小程序后台管理路由接口
		*/
		admin := auth.Group("/admin")
		{
			//eg.GET("/helloworld",Helloworld)
			//注册用户并获取token
			admin.GET("/user/Login", v1.Login)
			admin.GET("/getRefreshToken", v1.RefreshToken)
			//发送验证码
			admin.POST("/user/SendVerifyCode", v1.SendVerifyCode)
			//jwt授权的接口
			admin.Use(middleware.JwtToken())
			{
				// 用户模块的路由接口
				admin.GET("/user/getUserAndRoleInfoList", v1.GetUserAndRoleInfoList)
				admin.GET("/user/getUser", v1.GetSingleUserInfo)
				admin.GET("/user/getUserList", v1.GetUserList)
				admin.POST("/user/add", v1.AddUser)
				admin.PUT("/user/edit/:id", v1.EditUserInfo)
				admin.DELETE("/user/del/:id", v1.DeleteUserInfo)
				//修改密码
				/*admin.PUT("/user/changepw/:id", v1.ChangeUserPassword)*/
			}
		}
		/*
		  小程序前台调用的路由接口
		*/
		wechat := router.Group("/WeChatFront")
		{
			wechat.Use(middleware.JwtToken())
			{
				//小程序用户信息模块
				/*wechat.POST("user/add", v1.AddUser)
				wechat.GET("user/:id", v1.GetUserInfo)
				wechat.GET("users", v1.GetUsers)*/
			}
		}
	}
	_ = router.Run(common.HttpPort)
}
