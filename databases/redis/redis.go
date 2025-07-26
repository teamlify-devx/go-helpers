package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	cfg "github.com/spf13/viper"
	"time"
)

// NewRedisClient Returns new redis client
func NewRedisClient(ctx context.Context, dbNum int) (db *redis.Client, err error) {

	redisHost := fmt.Sprintf("%s:%d", cfg.GetString("Redis.HOST"), cfg.GetInt("Redis.PORT"))

	db = redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Username:     cfg.GetString("Redis.USER"),
		Password:     cfg.GetString("Redis.PASS"),
		MinIdleConns: cfg.GetInt("Redis.MIN_IDLE_CONN"),
		PoolSize:     cfg.GetInt("Redis.POOL_SIZE"),
		PoolTimeout:  time.Duration(cfg.GetInt("Redis.POOL_TIMEOUT")),
		DB:           dbNum,
	})

	_, err = db.Ping(ctx).Result()

	return
}
