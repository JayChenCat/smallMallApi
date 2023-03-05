package viewModels

import "time"

type UserInfo struct {
	Id       int64     `json:"Uid"`
	RoleId   int64     `json:"rid"`
	UserName string    `json:"userName"` //
	Mobile   string    `json:"mobile"`
	Email    string    `json:"email"`
	AddTime  time.Time `json:"createTime"` //
}
