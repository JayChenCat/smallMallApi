package models

import (
	"SmallMall/utils/common"
	"SmallMall/utils/erromsg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// 用户信息对象
type User struct {
	Id              int64     `gorm:"primarykey;column:Id;type:int"`
	RoleId          int64     `gorm:"column:RoleId;type:int"`              //
	UserName        string    `gorm:"column:UserName;type:varchar(20)"`    //
	Password        string    `gorm:"column:Password;type:varchar(100)"`   //
	Avatar          string    `gorm:"column:Avatar;type:varchar(20)"`      //
	NickName        string    `gorm:"column:NickName;type:varchar(20)"`    //
	Mobile          string    `gorm:"column:Mobile;type:varchar(20)"`      //
	Email           string    `gorm:"column:Email;type:varchar(20)"`       //
	LoginCount      int64     `gorm:"column:LoginCount;type:int"`          //
	LoginLastIp     string    `gorm:"column:LoginLastIp;type:varchar(20)"` //
	LoginLastTime   time.Time `gorm:"column:LoginLastTime"`                //
	IsLock          int       `gorm:"column:IsLock;type:int"`              //
	QRCode          string    `gorm:"column:QRCode;type:varchar(20)"`      //
	Address         string    `gorm:"column:Address;type:varchar(20)"`     //
	DepartID        int64     `gorm:"column:DepartID;type:int"`            //
	Token           string    `gorm:"column:Token;type:varchar(20)"`       //
	AddManagerId    int64     `gorm:"column:AddManagerId;type:int"`        //
	AddTime         time.Time `gorm:"column:AddTime"`                      //
	ModifyManagerId int64     `gorm:"column:ModifyManagerId;type:int"`     //
	ModifyTime      time.Time `gorm:"column:ModifyTime"`                   //
	IsDeleted       int       `gorm:"column:IsDeleted;type:int"`           //
	Remark          string    `gorm:"column:Remark;type:varchar(20)"`      //
}

/*func init() {
	user := User{
		Id:       18876868780,
		RoleId:   18876868785,
		UserName: "admin",
		Password: "$2a$10$YGL5a9z7ykG6BWOo.XhJU.h8r98BD5IvAmLISBB9rFIefbDzrv58O",
	}
	CreateUser(user)
}*/

//检查用户名是否存在
func CheckByUserName(name string) int {
	user := &User{}
	db.Select("id").Where("UserName=?", name).First(user)
	if user.Id > 0 {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpdateUser(id int64, name string, fields []string) (code int) {
	var user User
	db.Select(fields).Where("username = ?", name).First(&user)
	if user.Id == id {
		return erromsg.SUCCSE
	}
	if user.Id > 0 {
		return erromsg.ERROR_USERNAME_EXIST //1001
	}
	return erromsg.SUCCSE
}

// GetUser 查询用户
func GetSingleUser(id int64) (User, int) {
	var user User
	fields := []string{
		"Id", "RoleId", "UserName", "Password", "Avatar", "NickName",
		"Mobile", "Email", "LoginCount", "LoginLastIp",
		"LoginLastTime", "IsLock", "QRCode", "Address",
		"DepartID", "Token", "AddManagerId", "AddTime",
		"ModifyManagerId", "ModifyTime", "IsDeleted", "Remark",
	}
	conditions := map[string]interface{}{
		"Id": id,
	}
	err := db.Table("user").Select(fields).Where(conditions).Scan(&user).Error
	//err := db.First(&user, id).Error
	//sql := "select Id,RoleId,UserName,Mobile,Email,AddTime from user Where Id=?;"
	//err := db.Raw(sql, id).Scan(&user).Error
	if err != nil {
		common.WriteLog(err.Error())
		return user, erromsg.ERRORS
	}
	return user, erromsg.SUCCSE
}

//查询邮箱是否注册
func CheckEmail(email string) int {
	var num int64
	var user User
	err := db.Model(&user).Where("Email=?", email).Count(&num)
	if err != nil {
		return erromsg.ERRORS
	}
	if num > 0 {
		return erromsg.ERROR_EMAIL_EXIST
	}
	return erromsg.SUCCSE
}

// GetUsers 查询用户列表-分页
func GetPageUsers(pageSize int, pageNum int, fields []string,
	conditions map[string]interface{}) ([]User, int64) {
	var users []User
	var total int64
	//db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	offset := (pageNum - 1) * pageSize
	if len(conditions) > 0 {
		db.Select(fields).Where(conditions).Limit(pageSize).Offset(offset).Find(&users)
		db.Model(&users).Where(conditions).Count(&total)
		return users, total
	}
	db.Select(fields).Limit(pageSize).Offset(offset).Find(&users)
	db.Model(&users).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}

// CreateUser--新增用户
func CreateUser(user User) int {
	//data.Password = ScryptPw(data.Password)
	result := db.Table("user").Create(&user)
	err := result.Error
	if err != nil {
		return erromsg.ERRORS // 500
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return erromsg.SUCCSE
	}
	return erromsg.SUCCSE
}

//批量创建
func CreateBatches(value interface{}, batchSize int) int {
	//userList, len(userList)
	err := db.CreateInBatches(value, batchSize).Error
	if err != nil {
		return erromsg.ERRORS // 500
	}
	return erromsg.SUCCSE
}

// EditUser 编辑用户信息
func EditUser(id int64, columns map[string]interface{}) int {
	var user User
	/*var maps = make(map[string]interface{})
	maps["UserName"] = u.UserName
	maps["Email"] = u.Email*/
	result := db.Model(&user).Where("id = ? ", id).Updates(columns)
	err = result.Error
	if err != nil {
		return erromsg.ERRORS
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return erromsg.SUCCSE
	}
	return erromsg.SUCCSE
}

//批量更新
func BatchUpdate(query string, ids []int,
	columns map[string]interface{}) int {
	//"id IN (?)"
	//[]int{10, 11}
	//map[string]interface{}{"name": "hello", "age": 18}
	var user User
	err := db.Model(&user).Where(query, ids).Updates(columns).Error
	if err != nil {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int64) int {
	var user User
	err = db.Table("user").Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}

//删除记录(参数必须为结构体指针)-写法一
func DeleteFirst(key string, args int64, user *User) int {
	result := db.Table("user").Where(key, args).Delete(user) //"id =?"
	errs := result.Error                                     //返回 error
	if errs != nil {
		//log.Printf("删除表记录错误 error: %v", errs)
		common.WriteLog(err.Error())
		return erromsg.ERRORS
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return erromsg.SUCCSE
	}
	return erromsg.ERRORS
}

//批量删除-写法二(根据多个主键删除)
//批量更新
//IN关键字
//批量更新的话，只能批量对行的某些字段改为相同的值，不能改为不同的值…感觉没啥用
//db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})

func BatchDeleteDataByPrimarykey(key string, args []int64, user *User) int {
	result := db.Table("user").Where(key, args).Delete(user)
	errs := result.Error //返回 error
	if errs != nil {
		//log.Printf("删除表记录错误 error: %v", errs)
		common.WriteLog(err.Error())
		return erromsg.ERRORS
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return erromsg.SUCCSE
	}
	return erromsg.ERRORS
}

//批量删除调用
func DeleteById(ids []int64) int {
	var user User
	count := len(ids)
	if count == 1 {
		//单个删除
		//return DeleteFirst("Id=?", ids[0], user)
		err = db.Where("id = ? ", ids[0]).Delete(&user).Error
		if err != nil {
			common.WriteLog(err.Error())
			return erromsg.ERRORS
		}
		return erromsg.SUCCSE
	}
	if count > 1 {
		//批量删除
		return BatchDeleteDataByPrimarykey("Id in(?)", ids, &user)
	}
	return erromsg.ERRORS
}

// CheckLogin 小程序后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	//var PasswordErr error
	db.Select("Id,RoleId,UserName,Email").Where("UserName = ? and Password=?", username, password).First(&user) //Table("user")
	//PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Id == 0 {
		return user, erromsg.ERROR_USERNAME_NOT_EXIST
	}
	//if PasswordErr != nil {
	//return user, erromsg.ERROR_PASSWORD
	//}
	if user.RoleId == 0 {
		return user, erromsg.ERROR_USER_NO_AUTH
	}
	return user, erromsg.SUCCSE
}

// ScryptPassword 生成密码
func ScryptPassword(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

// ChangePassword 修改密码
func ChangePassword(id int64, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}
