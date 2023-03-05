package v1

import (
	"SmallMall/models"
	"SmallMall/service"
	_ "SmallMall/service"
	"SmallMall/utils/erromsg"
	"SmallMall/utils/middleware"
	"SmallMall/utils/validator"
	"SmallMall/viewModels"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// swagger:operation POST /user/GetSingleUserInfo
// @Summary 查询单个用户
// @Description 用于系统用户的查询
// @Tags 测试
// @Accept json
// @Param id query string true "用户ID(主键)"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/getUser  [get]
// @Security ApiKeyAuth
func GetSingleUserInfo(c *gin.Context) {
	userId := c.Query("id")
	id, _ := strconv.ParseInt(userId, 0, 0)
	data, code := models.GetSingleUser(id)
	/*var maps = make(map[string]interface{})
	maps["UserName"] = data.UserName
	maps["Email"] = data.Email*/
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": erromsg.GetErrosMsg(code),
		},
	)
}

// swagger:operation POST /user/GetUserList
// @Summary 查询用户列表
// @Description 用于系统用户的查询
// @Tags 测试
// @Accept json
// @Param pagenum query string true "当前页"
// @Param pagesize query string true "页码"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/getUserList   [get]
// @Security ApiKeyAuth
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	fields := []string{
		"Id",
		"RoleId",
		"UserName",
		"Password",
		"Avatar",
		"NickName",
		"Mobile",
		"Email",
		"LoginCount",
		"LoginLastIp",
		"LoginLastTime",
		"IsLock",
		"QRCode",
		"Address",
		"DepartID",
		"Token",
		"AddManagerId",
		"AddTime",
		"ModifyManagerId",
		"ModifyTime",
		"IsDeleted",
		"Remark",
	}
	conditions := map[string]interface{}{}
	if username != "" {
		conditions = map[string]interface{}{
			"UserName": username,
		}
	}
	data, total := models.GetPageUsers(pageSize, pageNum, fields, conditions)
	code := erromsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"total":   total,
			"message": erromsg.GetErrosMsg(code),
			"data":    data,
		},
	)
}

// swagger:operation POST /user/GetUserList
// @Summary 查询用户列表(2)
// @Description 用于系统用户的查询
// @Tags 测试
// @Accept json
// @Param username query string false "查询-用户名称"
// @Param pagenum query string true "当前页"
// @Param pagesize query string true "页码"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/getUserAndRoleInfoList  [get]
func GetUserAndRoleInfoList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	fields := []string{
		"user.Id as UId",
		"RoleId  as RId",
		"RoleName",
		"UserName",
		"user.AddTime as CreateTime",
	}
	conditions := map[string]interface{}{}
	if username != "" {
		conditions = map[string]interface{}{
			"UserName": username,
		}
	}
	data, total := service.GetUserAndRoleInfoList(pageSize, pageNum, fields, conditions)
	code := erromsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": erromsg.GetErrosMsg(code),
		},
	)
}

// swagger:operation POST /user/addUser
// @Summary 新增用户信息
// @Description 用于系统用户的新增
// @Tags 测试
// @Accept json
// @Param  User body models.User true "提交用户信息"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/add  [post]
// @Security ApiKeyAuth
func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	msg, validCode := validator.ValidateParam(&data)
	if validCode != erromsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}
	code := models.CheckByUserName(data.UserName)
	if code == erromsg.SUCCSE {
		code = models.CreateUser(data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": erromsg.GetErrosMsg(code),
		},
	)
}

// swagger:operation POST /user/EditUserInfo
// @Summary 编辑用户信息
// @Description 用于编辑系统用户信息
// @Tags 测试
// @Accept json
// @Param  User body models.User true "提交用户信息"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/edit/{id}  [put]
// @Security ApiKeyAuth
func EditUserInfo(c *gin.Context) {
	var data models.User
	/*userId := c.Query("id")
	id, _ := strconv.ParseInt(userId, 0, 0)*/
	_ = c.ShouldBindJSON(&data)
	id := data.Id
	fields := []string{
		"Id",
		"UserName",
		"Avatar",
		"NickName",
		"Mobile",
		"Email",
	}
	code := models.CheckUpdateUser(id, data.UserName, fields)
	if code == erromsg.SUCCSE {
		var maps = make(map[string]interface{})
		maps["UserName"] = data.UserName
		maps["Email"] = data.Email
		code = models.EditUser(id, maps)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": erromsg.GetErrosMsg(code),
		},
	)
}

// swagger:operation POST /user/DeleteUserInfo
// @Summary 删除用户
// @Description 用于删除系统单个信息
// @Tags 测试
// @Accept json
// @Param id query string true "用户ID(主键)"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/del/{id}  [delete]
// @Security ApiKeyAuth
func DeleteUserInfo(c *gin.Context) {
	userId := c.Query("id")
	id, _ := strconv.ParseInt(userId, 0, 0)
	ids := []int64{id}
	code := models.DeleteById(ids)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": erromsg.GetErrosMsg(code),
		},
	)
}

// swagger:operation POST /admin/Login
// @Summary 后台登陆获取token
// @Description 用于用户登录系统管理后台
// @Tags 测试
// @Accept json
// @Param UserName query   string true "用户名"
// @Param Password query   string true "密码"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/Login   [get]
func Login(c *gin.Context) {
	formData := models.User{}
	userName := c.Query("UserName")
	password := c.Query("Password")
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int
	formData, code = models.CheckLogin(userName, password)
	if code == erromsg.SUCCSE {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.UserName,
			"id":      formData.Id,
			"message": erromsg.GetErrosMsg(code),
			"token":   token,
		})
	}
}

// swagger:operation POST
// @Summary 发送验证码
// @Description 向用户绑定的邮箱发送验证码
// @Tags 验证码测试
// @Accept json
// @Param  MailBoxInfo body viewModels.MailBoxViewModel true "提交用户信息"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/user/SendVerifyCode   [post]
func SendVerifyCode(c *gin.Context) {
	formData := viewModels.MailBoxViewModel{}
	_ = c.ShouldBindJSON(&formData)
	email := formData.Email //chensj@sqray.com
	uid := formData.UID
	/*num, errs := strconv.ParseInt(uid, 0, 0)
	if errs != nil {
		log.Printf("uid转换整型的错误 error: %v", errs.Error())
	}*/
	recipientsList := []string{email}
	code := service.SendEmailCode(uid, recipientsList)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": erromsg.GetErrosMsg(code),
	})
	return
}

// swagger:operation POST /admin/getRefreshToken
// @Summary 获取更新的token
// @Description 用于用户登录系统管理后台
// @Tags 测试
// @Accept json
// @Param token query   string true "旧的token"
// @Success 200 {string} string "{"status": 200,"message":"成功"}"
// @Failure 400 {string} string "{"status": 500,"message":"失败"}"
// @Router /admin/getRefreshToken   [get]
func RefreshToken(c *gin.Context) {
	_token := c.Query("token")
	j := middleware.NewJWT()
	//_ = c.ShouldBindJSON(&formData)
	token, err := j.RefreshToken(_token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  erromsg.ERRORS,
			"message": erromsg.GetErrosMsg(erromsg.ERRORS),
			"token":   token,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": erromsg.GetErrosMsg(200),
		"token":   token,
	})
}

// token生成函数
func setToken(c *gin.Context, user models.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,     // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, // 过期时间 2h
			Issuer:    "CSJ",                       // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  erromsg.ERRORS,
			"message": erromsg.GetErrosMsg(erromsg.ERRORS),
			"token":   token,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.UserName,
		"id":      user.Id,
		"message": erromsg.GetErrosMsg(200),
		"token":   token,
	})
	return
}
