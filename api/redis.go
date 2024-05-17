package api

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = &redis.Client{}

// InitRedisClient NewRedisClient Initialize Redis client
func InitRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// LoadStock Load initial stock value

func LoadStock(itemID string, stock int) error {
	return rdb.Set(ctx, itemID, stock, 0).Err()
}

func GetStock(itemID string) (int, error) {
	return rdb.Get(ctx, itemID).Int()
}
