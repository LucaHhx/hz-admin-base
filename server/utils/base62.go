package utils

import (
	"strings"
)

const (
	codeLength  = 8
	base36Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base        = 36
	maxNumber   = 2821109907456 // 36^8
	prime       = 899809343     // 任意大质数，和maxNumber互质
)

// 混淆ID
func obfuscate(id uint) uint {
	return (id * prime) % maxNumber
}

// Base36编码
func encodeBase36(n uint) string {
	var encoded strings.Builder
	for n > 0 {
		remainder := n % base
		encoded.WriteByte(base36Chars[remainder])
		n /= base
	}

	// 倒序
	code := encoded.String()
	runes := []rune(code)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 补0
	padded := string(runes)
	for len(padded) < codeLength {
		padded = "0" + padded
	}
	return padded
}

// 生成唯一码
func GenerateCode(id uint) string {
	return encodeBase36(obfuscate(id))
}
