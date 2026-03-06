package utils

import (
	crand "crypto/rand"
	"math/big"
	"strings"
)

func RandNum6() string {
	bigInt, _ := crand.Int(crand.Reader, big.NewInt(900000))
	bigInt.Add(bigInt, big.NewInt(100000))
	return bigInt.String()
}

// RandInt 随机一个值
func RandInt(v int) int {
	if v <= 0 {
		return 0
	}
	bigInt, _ := crand.Int(crand.Reader, big.NewInt(int64(v)))
	return int(bigInt.Int64())
}

func RandString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder

	// 确保字符串长度有效
	if length <= 0 {
		length = 16
	}

	// 生成随机字符
	for i := 0; i < length; i++ {
		randChar := chars[RandInt(len(chars))]
		sb.WriteByte(randChar)
	}
	return sb.String()
}
