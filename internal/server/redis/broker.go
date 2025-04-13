package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func PublishMessage(msg string) {
	err := rdb.Publish(ctx, "soniq-messages", msg).Err()
	if err != nil {
		fmt.Println("Redis publish error:", err)
	}
}

func Subscribe(callback func(string)) {
	sub := rdb.Subscribe(ctx, "soniq-messages")
	ch := sub.Channel()

	go func() {
		for msg := range ch {
			callback(msg.Payload)
		}
	}()
}
