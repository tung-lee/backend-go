package initialization

import (
	"backend-go/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	redisSetting := global.Config.Redis
	
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%v", redisSetting.Host, redisSetting.Port),
        Password: redisSetting.Password, // no password set
        DB:       redisSetting.Database,  // use default DB
		PoolSize: 10,
    })

	_, err := rdb.Ping(ctx).Result() // Check if Redis is running

	if err != nil {
		global.Logger.Error("Redis initialization failed: %v", zap.Error(err))
		panic(err)
	}

	fmt.Println("Redis initialization success")
	global.RedisDB = rdb
	redisExample()
}

func redisExample() {
	err := global.RedisDB.Set(ctx, "point", 100, 0).Err()

	if err != nil {
		fmt.Println("Redis set error: ", err)
		return
	}

	val, err := global.RedisDB.Get(ctx, "point").Result()

	if err != nil {
		fmt.Println("Redis get error: ", err)
		return
	}

	global.Logger.Info("Redis get success - value point: ", zap.String("point", val))
}