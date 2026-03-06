package utils

import (
	"fmt"
	"github.com/samber/lo"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomIntString(n int) string {
	var letters = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

//func RandUuid(key, begin string, startLong int) string {
//	for i := 0; i < 10; i++ {
//		startLong = startLong + i
//		for i := 0; i < 10; i++ {
//			nowkey := begin + RandomString(startLong)
//			result, err := global.HAB_REDIS.SetNX(context.Background(), fmt.Sprintf("%s:{%s}:%s", enum.NxUid, key, nowkey), 1, 24*time.Hour).Result()
//			if err != nil || !result {
//				continue
//			}
//			return nowkey
//		}
//	}
//	return ""
//}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func ArrayConv[T, V any](ts []T, f func(t T) (v V, ok bool)) []V {
	vs := make([]V, 0)
	for _, t := range ts {
		v, ok := f(t)
		if ok {
			vs = append(vs, v)
		}
	}
	return vs
}

func ArrJoin[T any](arr []T, sep string) string {
	var strs []string
	for _, v := range arr {
		strs = append(strs, fmt.Sprint(v))
	}
	return strings.Join(strs, sep)
}

func ArrSplitInt(str, sep string) []int {
	var arr []int
	for _, v := range strings.Split(str, sep) {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		arr = append(arr, atoi)
	}
	return arr
}

func Split(str, sep string) []string {
	strs := lo.Filter(strings.Split(str, sep), func(item string, index int) bool {
		return item != ""
	})
	var arr []string
	for _, v := range strs {
		arr = append(arr, strings.TrimSpace(v))
	}
	return arr
}

func SplitNum[T Number](str, sep string) []T {
	var arr []T
	for _, v := range strings.Split(str, sep) {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		arr = append(arr, T(atoi))
	}
	return arr
}
