package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

// redis客户端
var rdb *redis.Client

// initRedis
//
//	@Description: 初始化redis
//	@return error error
func initRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "10.40.3.29:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Printf("redis ping error: %s\n", err)
		return err
	}

	return nil
}

// 判断redis中key存在
// 如果返回的n == 0则说明key不存在
func ExistKey(key string) bool {
	//执行redis命令Exists，返回key的存在状态
	n, err := rdb.Exists(key).Result()
	if err != nil {
		fmt.Printf("redis Exists error: %s\n", err)
		return false
	}

	//如果key不存在，则返回n == 0，如果key存在，则返回n > 0
	if n == 0 {
		fmt.Printf("key %s not exists\n", key)
		return false
	}

	fmt.Printf("key %s exists\n", key)
	return true
}

// 设置或获取缓存
// 根据文章Id获取缓存
func GetCachePostById(key string) {
	result, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Printf("redis Get post error: %s\n", err)
		return
	}

	json.Unmarshal([]byte(result))
}
