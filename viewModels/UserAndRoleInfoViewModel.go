package viewModels

import "time"

type UserAndRoleInfoList struct {
	UId        int64     `json:"Uid"`
	RId        int64     `json:"rid"`
	RoleName   string    `json:"roleName"`   //
	UserName   string    `json:"userName"`   //
	CreateTime time.Time `json:"createTime"` //
}
