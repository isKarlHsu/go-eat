package core

import (
	"context"
	"eat/global"
	"github.com/go-redis/redis"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDb(global.Config.Redis.Db)
}

func ConnectRedisDb(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Logger.Errorf("redis连接失败 %s", err)
		return nil
	}
	return rdb
}
