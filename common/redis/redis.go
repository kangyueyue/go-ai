package redis

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/kangyueyue/go-ai/config"
)

var Rdb *redis.Client

func Init() {
	conf := config.GetConfig()
	host := conf.RedisConfig.RedisHost
	port := conf.RedisConfig.RedisPort
	password := conf.RedisConfig.RedisPassword
	db := conf.RedisDb
	addr := host + ":" + strconv.Itoa(port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// CheckCaptcha 校验验证码
func CheckCaptcha(email, userInputCaptcha string) (bool, error) {
	key := GenerateCaptcha(email)

	captcha, err := Rdb.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			// redis不存在
			return false, nil
		}
		return false, err
	}
	if strings.EqualFold(captcha, userInputCaptcha) {
		// 验证成功，del
		if err := Rdb.Del(context.Background(), key).Err(); err != nil {
			// 删除失败，不影响主要注册
		}
		return true, nil
	}
	// 验证码错误
	return false, nil
}
