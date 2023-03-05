package goredis

import (
	"SmallMall/utils/common"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//操作Redis-普通连接
func redisinit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "119.91.142.127:6379",
		Password: "123456",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		common.WriteLog(fmt.Sprintf("Redis-普通连接错误:%s", err.Error()))
	}
	return rdb
}

//向redis存入值-时间单位:分钟
func Set(key string, value interface{}, expiration time.Duration) bool {
	rdb := redisinit()
	//设置一分钟的有效期
	//rdb.Expire(ctx, "key", time.Minute)
	err := rdb.Set(key, value, expiration).Err()
	if err != nil {
		common.WriteLog(fmt.Sprintf("向redis存入值错误:%s", err.Error()))
		return false
	}
	return true
}

//通过key从redis中读取值
func GetValue(key string) string {
	rdb := redisinit()
	val, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return ""
	}
	if err != nil {
		common.WriteLog(fmt.Sprintf("通过key从redis中读取值错误:%s", err.Error()))
		return ""
	}
	return val
}

//通过key删除值
func DeleteValue(key string) bool {
	rdb := redisinit()
	_, err := rdb.Del(key).Result()
	if err != nil {
		common.WriteLog(fmt.Sprintf("通过key删除值错误:%s", err.Error()))
		return false
	}
	return true
}

//检测缓存项是否存在
func Exists(key string) bool {
	rdb := redisinit()
	n, err := rdb.Exists(key).Result()
	if err != nil {
		common.WriteLog(fmt.Sprintf("通过key检测缓存项是否存在错误:%s", err.Error()))
	}
	if n > 0 {
		return true
	}
	return false
}

//获取存入的键值有效期-时间单位:分钟
func GetExpire(key string) float64 {
	rdb := redisinit()
	ttl, err := rdb.TTL(key).Result()
	if err != nil {
		common.WriteLog(fmt.Sprintf("获取存入的键值有效期错误:%s", err.Error()))
	}
	return ttl.Minutes()
}
