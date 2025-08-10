package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	cfg "github.com/spf13/viper"
)

// NewDFClient Returns new DragonFly (using same driver with redis) client
func NewDFClient(ctx context.Context, dbNum int) (db *redis.Client, err error) {

	connStr := fmt.Sprintf("%s:%d", cfg.GetString("Dragonfly.HOST"), cfg.GetInt("Dragonfly.PORT"))

	db = redis.NewClient(&redis.Options{
		Addr:         connStr,
		Username:     cfg.GetString("Dragonfly.USER"),
		Password:     cfg.GetString("Dragonfly.PASS"),
		MinIdleConns: cfg.GetInt("Dragonfly.MIN_IDLE_CONN"),
		PoolSize:     cfg.GetInt("Dragonfly.POOL_SIZE"),
		PoolTimeout:  time.Duration(cfg.GetInt("Dragonfly.POOL_TIMEOUT")),
		DB:           dbNum,
	})

	_, err = db.Ping(ctx).Result()

	return
}
