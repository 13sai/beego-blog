package services

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"encoding/json"
	"fmt"
	"time"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", beego.AppConfig.String("redis_host"), beego.AppConfig.String("redis_port")),
		Password: beego.AppConfig.String("redis_auth"), // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}

func SetCache(key string, val interface{}) {
	result, _ := json.Marshal(val)
	_, err := RedisClient.Set(key, result, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func SetCacheWT(key string, val interface{}, ttl int) {
	result, _ := json.Marshal(val)
	_, err := RedisClient.Set(key, result, time.Second*time.Duration(ttl)).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func GetCache(key string) (ret interface{}) {
	if RedisClient.Exists(key).Val() > 0 {
		by, err := RedisClient.Get(key).Bytes();
		if err !=nil {
			fmt.Println("error")
		}

		json.Unmarshal(by, &ret)
		return ret
	}

	return nil
}