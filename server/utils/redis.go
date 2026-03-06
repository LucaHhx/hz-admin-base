package utils

import (
	"context"
	"fmt"
	"hab/enum"
	"hab/global"

	"github.com/samber/lo"
)

// LuaHINCRBY 验证增量操作不能小于0
func LuaHINCRBY(key string, field string, value int64) (int64, error) {
	luaScript := `
		local current = tonumber(redis.call("HGET", KEYS[1], ARGV[1])) or 0
		local delta = tonumber(ARGV[2])

		-- 如果是减少金额，检查是否足够
		if delta < 0 then
			if current + delta < 0 then
				return -1 -- 余额不足
			end
		end

		-- 执行 HINCRBY
		return redis.call("HINCRBY", KEYS[1], ARGV[1], delta)
	`
	return global.HAB_REDIS.Eval(context.Background(), luaScript, []string{key}, field, value).Int64()
}

func LuaHINCRBYList(key string, fields []string, values []int64) ([]int64, error) {
	luaScript := StrReplace(`
		local hashKey = "@hashKey"
		local results = {}

		-- 先检查所有字段是否满足扣减条件
		for i = 1, #KEYS do
			local field = KEYS[i]
			local delta = tonumber(ARGV[i])
			local current = tonumber(redis.call("HGET", hashKey, field)) or 0

			if delta < 0 and current + delta < 0 then
				return {-1} -- 如果有余额不足，直接返回错误
			end
		end

		-- 如果所有字段都满足条件，则执行 HINCRBY
		for i = 1, #KEYS do
			local field = KEYS[i]
			local delta = tonumber(ARGV[i])
			table.insert(results, redis.call("HINCRBY", hashKey, field, delta))
		end

		return results
	`, map[string]string{
		"hashKey": key,
	})
	args := make([]interface{}, 0)
	for _, value := range values {
		args = append(args, value)
	}
	slice, err := global.HAB_REDIS.Eval(context.Background(), luaScript, fields, args...).Int64Slice()
	if err != nil {
		return nil, enum.Msg_InternalError
	}
	if len(lo.Filter(slice, func(i int64, index int) bool { return i == -1 })) > 0 {
		return nil, enum.Msg_MerchantInsufficient
	}
	return slice, nil
}

func LuaHGETInt64(key string, fields ...string) ([]int64, error) {
	luaScript := fmt.Sprintf(`
		local results = {}
		for i, field in ipairs(KEYS) do
			table.insert(results, redis.call("HGET", "%s" , KEYS[i]) or 0)
		end
		return results
	`, key)
	args := make([]interface{}, 0)
	args = append(args, fields)
	return global.HAB_REDIS.Eval(context.Background(), luaScript, fields).Int64Slice()
}
