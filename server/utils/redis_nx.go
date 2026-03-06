package utils

import (
	"context"
	"hz-admin-base/global"
	"time"
)

type RedisNx struct {
	key string
}

func NewNx(key string) *RedisNx {
	redisNx := &RedisNx{key: key}
	return redisNx
}

func NewRedisNx(key string, expiration time.Duration) *RedisNx {
	redisNx := &RedisNx{key: key}
	ok := redisNx.SetNx(expiration)
	if !ok {
		return nil
	}
	return redisNx
}

func DeleteRedisNx(key string) {
	redisNx := &RedisNx{key: key}
	redisNx.Del()
}

func (r *RedisNx) SetNx(expiration time.Duration) bool {
	result, _ := global.GVA_REDIS.SetNX(context.Background(), r.key, 1, expiration).Result()
	return result
}

func (r *RedisNx) Del() {
	global.GVA_REDIS.Del(context.Background(), r.key)
}

func (r *RedisNx) IsLocked() bool {
	val, err := global.GVA_REDIS.Get(context.Background(), r.key).Result()
	if err != nil {
		return false
	}
	return val == "1"
}
