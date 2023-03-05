package viewModels

import "time"

type CodeInfoViewModel struct {
	Id      int64     `json:"id"`
	UId     int64     `json:"uid"`
	Code    string    `json:"code"`    //验证码
	Expire  float64   `json:"expire"`  //有效期，默认3分钟
	AddTime time.Time `json:"addTime"` //验证码添加时间
}
