package common

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

// 记录日志
// s interface{} 日志内容
// return error
func WriteLog(s interface{}) error {
	filePath := "./log/" + time.Now().Format("20060102") + "-error.txt" //150405
	path, _ := filepath.Split(filePath)                                 // 获取路径
	_, err := os.Stat(path)                                             // 检查路径状态，不存在创建
	if err != nil || os.IsExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	fp, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666) // 读或创建形式打开文件
	defer fp.Close()
	b, err := json.Marshal(s) // 将日志转Json
	if err != nil {
		return errors.New(err.Error())
	}
	b = append(b, 10)              // 添加换行
	n, _ := fp.Seek(0, io.SeekEnd) // 偏移量位置
	_, err = fp.WriteAt(b, n)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
