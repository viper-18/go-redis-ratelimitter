package app

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimitter struct {
	client     *redis.Client
	rate       int
	windowSize time.Duration
}

func NewRateLimitter(client *redis.Client, rate int, windowSize time.Duration) *RateLimitter {
	return &RateLimitter{
		client:     client,
		rate:       rate,
		windowSize: windowSize,
	}
}

func (rl *RateLimitter) AllowRequest(clientID string) (bool, error) {
	key := fmt.Sprintf("rate_limit:%s", clientID)
	now := time.Now().UnixNano()

	ctx := context.Background()

	pipe := rl.client.TxPipeline()

	// Queue Redis commands in the pipeline
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})
	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", now-int64(rl.windowSize)))
	pipe.Expire(ctx, key, rl.windowSize*2)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	count := rl.client.ZCard(ctx, key).Val()
	if count > int64(rl.rate) {
		return false, nil
	}

	return true, nil
}
