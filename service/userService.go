package service

import (
	"SmallMall/models"
	"SmallMall/viewModels"
)

func GetUserAndRoleInfoList(pageSize int, pageNum int, fields []string, conditions map[string]interface{}) ([]viewModels.UserAndRoleInfoList, int64) {
	lists, code := models.GetUserAndRoleInfoList(pageSize, pageNum, fields, conditions)
	return lists, code
}
