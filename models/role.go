package models

import (
	"SmallMall/viewModels"
	"time"
)

type Role struct {
	Id              int64     `gorm:"column:Id;primarykey"`
	RoleName        string    `gorm:"column:RoleName;type:varchar(20)"` //
	AddManagerId    int64     `gorm:"column:AddManagerId;type:int"`     //
	AddTime         time.Time `gorm:"column:AddTime"`                   //
	ModifyManagerId int64     `gorm:"column:ModifyManagerId;type:int"`  //
	ModifyTime      time.Time `gorm:"column:ModifyTime"`                //
	IsDeleted       int       `gorm:"column:IsDeleted;type:int"`        //
	Remark          string    `gorm:"column:Remark;type:varchar(20)"`   //
}

func GetUserAndRoleInfoList(pageSize int, pageNum int, fields []string, conditions map[string]interface{}) ([]viewModels.UserAndRoleInfoList, int64) {
	var users []viewModels.UserAndRoleInfoList
	joinsSQL := "left join role on role.Id =user.RoleId"
	tableName := "user"
	var total int64
	//db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	offset := (pageNum - 1) * pageSize
	if len(conditions) > 0 {
		db.Table(tableName).Select(fields).Joins(joinsSQL).Where(conditions).Limit(pageSize).Offset(offset).Scan(&users)
		db.Table(tableName).Select(fields).Joins(joinsSQL).Where(conditions).Count(&total)
		return users, total
	}
	/*sqlCount := "select count(1) as total from user u  left join role r on r.Id = u.RoleId;"
	sql := "select u.Id UId,RoleId  RId,RoleName, UserName, u.AddTime CreateTime from user u  left join role r on r.Id = u.RoleId  limit ? offset ?;"
	db.Raw(sql, pageSize, offset).Scan(&users)
	db.Raw(sqlCount).Scan(&total)
	//db.Raw(sqlCount).Scan(&total)*/
	db.Table(tableName).Select(fields).Joins(joinsSQL).Limit(pageSize).Offset(offset).Scan(&users)
	db.Table(tableName).Select(fields).Joins(joinsSQL).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}
