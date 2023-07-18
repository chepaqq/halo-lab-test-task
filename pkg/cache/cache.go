package cache

import "github.com/redis/go-redis/v9"

type Config struct {
	Addr     string
	Password string
	DB       int
}

// ConnectRedis - .
func ConnectRedis(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
