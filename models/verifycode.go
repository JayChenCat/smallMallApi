package models

import (
	"SmallMall/utils/common"
	"SmallMall/utils/erromsg"
	"time"
)

type Code struct {
	Id      int64     `gorm:"column:Id;primarykey";label:"主键ID"`           //主键
	UID     int64     `gorm:"column:UID;type:int";label:"关联的用户ID"`         //关联的用户ID,一个用户对应多个验证码
	Code    string    `gorm:"column:Code;type:varchar(50)";label:"验证码"`    //验证码
	Expire  float64   `gorm:"column:Expire;type:double";label:"有效期，默认3分钟"` //有效期，默认3分钟
	AddTime time.Time `gorm:"column:AddTime";label:"验证码添加时间"`              //验证码添加时间
}

//AddCode --添加验证码(存放验证码记录)
func AddCode(code Code) int {
	result := db.Table("code").Create(&code)
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

// DeleteCode 删除验证码记录-针对过期的验证码
func DeleteCode(id int64) int {
	var code Code
	err = db.Table("code").Where("id = ? ", id).Delete(&code).Error
	if err != nil {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}

// DeleteCode 删除验证码记录-针对过期的验证码
func ByUidDeleteCode(uid int64) int {
	var code Code
	err = db.Table("code").Where("UID=?", uid).Delete(&code).Error
	if err != nil {
		return erromsg.ERRORS
	}
	return erromsg.SUCCSE
}

var fields = []string{
	"Id", "UID", "Code", "Expire", "AddTime",
}

//根据存放验证码记录ID查询
func GetSingleCode(id int64) (Code, int) {
	var code Code
	conditions := map[string]interface{}{
		"Id": id,
	}
	err := db.Table("code").Select(fields).Where(conditions).Scan(&code).Error
	if err != nil {
		common.WriteLog(err.Error())
		return code, erromsg.ERRORS
	}
	return code, erromsg.SUCCSE
}

//根据验证码用户ID查询对应验证码记录
func ByUserIdGetCode(uid int64) (Code, int) {
	var code Code
	/*conditions := map[string]interface{}{
		"Uid": uid,
	}*/
	err := db.Table("code").Select(fields).Where("UID=?", uid).Find(&code).Error
	//已存在
	if err != nil {
		common.WriteLog(err.Error())
		return code, erromsg.ERRORS
	}
	if code.UID > 0 {
		return code, erromsg.ERROR_EMAIL_EXIST
	}
	return code, erromsg.ERROR_EMAIL_NOEXIST
}

// EditCode 编辑验证码记录
func EditCode(id int64, columns map[string]interface{}) int {
	var code Code
	/*var maps = make(map[string]interface{})
	maps["UserName"] = u.UserName
	maps["Email"] = u.Email*/
	result := db.Model(&code).Where("id = ? ", id).Updates(columns)
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
