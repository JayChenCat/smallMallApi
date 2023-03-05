package service

import (
	"SmallMall/models"
	"SmallMall/viewModels"
)

// EditCodeInfo--新增或修改验证码记录
func EditCodeInfo(code viewModels.CodeInfoViewModel) int {
	_code := models.Code{
		Id:      code.Id,
		UID:     code.UId,
		Code:    code.Code,
		Expire:  code.Expire,
		AddTime: code.AddTime,
	}
	//参与编辑的列与参数键值对
	maps := map[string]interface{}{
		"Uid":     code.UId,
		"Code":    code.Code,
		"Expire":  code.Expire,
		"AddTime": code.AddTime,
	}
	if code.Id > 0 {
		return models.EditCode(_code.Id, maps)
	} else {
		return models.AddCode(_code)
	}
}

//根据存放验证码记录ID查询
func GetSingleCodeInfo(id int64) (viewModels.CodeInfoViewModel, int) {
	code, num := models.GetSingleCode(id)
	_code := viewModels.CodeInfoViewModel{
		Id:      code.Id,
		UId:     code.UID,
		Code:    code.Code,
		Expire:  code.Expire,
		AddTime: code.AddTime,
	}
	return _code, num
}

//根据验证码用户ID查询对应验证码记录
func GetCode(uid int64) (viewModels.CodeInfoViewModel, int) {
	code, num := models.ByUserIdGetCode(uid)
	_code := viewModels.CodeInfoViewModel{
		Id:      code.Id,
		UId:     code.UID,
		Code:    code.Code,
		Expire:  code.Expire,
		AddTime: code.AddTime,
	}
	return _code, num
}

//根据用户ID删除验证码记录
func DeleteCodeInfo(uid int64) int {
	return models.ByUidDeleteCode(uid)
}
