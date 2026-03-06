package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"hab/global"
	"strconv"
	"strings"
	"time"
)

func GetCacheKey(place string, keys ...string) string {
	return place + ":" + strings.Join(keys, ":")
}

func GetCache(key string, to interface{}) error {
	res, err := global.HAB_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		//global.HAB_LOG.Error("GetCache Error: ", zap.Error(err))
		return err
	}
	return JsonDecode(res, to)
}

func SetCache(key string, data interface{}, expiration time.Duration) error {
	res, err := JsonEncode(data)
	if err != nil {
		global.HAB_LOG.Error("SetCache Error: ", zap.Error(err))
		return err
	}
	err = global.HAB_REDIS.Set(context.Background(), key, res, expiration).Err()
	if err != nil {
		global.HAB_LOG.Error("SetCache Error: ", zap.Error(err))
	}
	return err
}

func DelCache(key ...string) error {
	err := global.HAB_REDIS.Del(context.Background(), key...).Err()
	if err != nil {
		global.HAB_LOG.Error("DelCache Error: ", zap.Error(err))
	}
	return err
}

func SetHCache(key string, field string, data interface{}) error {
	res, err := JsonEncode(data)
	if err != nil {
		global.HAB_LOG.Error("SetHCache Error: ", zap.Error(err))
		return err
	}
	err = global.HAB_REDIS.HSet(context.Background(), key, field, res).Err()
	if err != nil {
		global.HAB_LOG.Error("SetHCache Error: ", zap.Error(err))
	}
	return err
}

func GetHCache(key string, field string, to interface{}) error {
	res, err := global.HAB_REDIS.HGet(context.Background(), key, field).Bytes()
	if err != nil {
		global.HAB_LOG.Error("GetHCache Error: ", zap.Error(err))
		return err
	}
	return JsonDecode(res, to)
}

func DelHCache(key string, field string) error {
	err := global.HAB_REDIS.HDel(context.Background(), key, field).Err()
	if err != nil {
		global.HAB_LOG.Error("DelHCache Error: ", zap.Error(err))
	}
	return err
}

func GetHAllCache[T any](key string, maps map[string]T) error {
	// 获取所有字段和值
	res, err := global.HAB_REDIS.HGetAll(context.Background(), key).Result()
	if err != nil {
		global.HAB_LOG.Error("GetHAllCache Error: ", zap.Error(err))
		return err
	}

	// 遍历所有字段和值，并解码到指定类型
	for field, value := range res {
		var item T
		if err = JsonDecode([]byte(value), &item); err != nil {
			global.HAB_LOG.Error("GetHAllCache Error: ", zap.Error(err))
			return err
		}
		maps[field] = item
	}
	return nil
}

type Eval struct {
	Key     string
	Operate string
	Args    []any
}

// AtomicOperate Eval原子操作
func AtomicOperate(evals ...*Eval) *redis.Cmd {
	var (
		builder strings.Builder
		argKey  = 1
		args    []any
		keys    []string
	)
	builder.WriteString("return {")
	for i, eval := range evals {
		builder.WriteString("redis.call('")
		builder.WriteString(eval.Operate)
		builder.WriteString("', KEYS[")
		builder.WriteString(strconv.Itoa(i + 1))
		builder.WriteString("]")
		// 拼接每个 eval 的参数
		for _, arg := range eval.Args {
			builder.WriteString(", ARGV[")
			builder.WriteString(strconv.Itoa(argKey))
			builder.WriteString("]")
			args = append(args, arg)
			argKey++
		}
		builder.WriteString("),")
		keys = append(keys, eval.Key)
	}
	builder.WriteString("}")
	command := builder.String()
	//fmt.Println(fmt.Sprintf(`eval "%s" %d %s %s`, command, len(keys), strings.Join(keys, " "), helper.AnyJoin(args, " ")))
	return global.HAB_REDIS.Eval(context.Background(), command, keys, args...)
}
