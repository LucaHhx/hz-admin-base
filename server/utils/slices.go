package utils

import (
	"strconv"
	"strings"
	"sync"
)

func ContainsIn[T comparable](ts []T, t T) bool {
	for _, item := range ts {
		if item == t {
			return true
		}
	}
	return false
}

func ContainsInArr[T comparable](ts []T, nts []T) bool {
	for _, item := range ts {
		for _, nItem := range nts {
			if item == nItem {
				return true
			}
		}
	}
	return false
}

func ArrIn[T comparable](ts []T, nts ...T) bool {
	for _, nt := range nts {
		if !ContainsIn(ts, nt) {
			return false
		}
	}
	return true
}

func ArrNotIn[T comparable](ts []T, nts ...T) bool {
	for _, nt := range nts {
		if ContainsIn(ts, nt) {
			return false
		}
	}
	return true
}

// SplitInt 字符串分割为整型数组
func SplitInt[V Int](s string, sep string) []V {
	arr := strings.Split(s, sep)
	var intArr []V
	for _, s := range arr {
		if s == "" {
			continue
		}
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, V(v))
	}
	return intArr
}

func JoinInt[V Int](arr []V, sep string) string {
	var strArr []string
	for _, s := range arr {
		strArr = append(strArr, strconv.Itoa(int(s)))
	}
	return strings.Join(strArr, sep)
}

func IntArr(arr []string) []int {
	var intArr []int
	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, v)
	}
	return intArr
}

func IntArrType[T, V Number](arr []V) []T {
	var intArr []T
	for _, s := range arr {
		intArr = append(intArr, T(s))
	}
	return intArr
}

func StringArr[V int | int64 | uint | uint64](arr []V) []string {
	var stringArr []string
	for _, s := range arr {
		stringArr = append(stringArr, strconv.FormatInt(int64(s), 10))
	}
	return stringArr
}

// InArr 某个值是否在数组中
func InArr[V comparable](v V, sl []V) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func SliceGet[T any](s []T, index int) T {
	// 转换负索引为正索引
	if index < 0 {
		index = len(s) + index
	}

	// 检查索引是否在有效范围内
	if index < 0 || index >= len(s) {
		// 返回T类型的零值
		var zeroValue T
		return zeroValue
	}

	return s[index]
}

func ArrFind[T any](s []T, f func(T) bool) *T {
	for _, v := range s {
		if f(v) {
			return &v
		}
	}
	return nil
}

// GetValues 获取指针数组中的值
func GetValues[T any](arr []*T) []T {
	var newArr []T
	for _, v := range arr {
		newArr = append(newArr, *v)
	}
	return newArr
}

type LockArr[T any] struct {
	arr []T
	mu  sync.Mutex
}

func NewLockArr[T any]() *LockArr[T] {
	return &LockArr[T]{arr: make([]T, 0)}
}

func (la *LockArr[T]) Add(vs ...T) {
	la.mu.Lock()
	defer la.mu.Unlock()
	la.arr = append(la.arr, vs...)
}

func (la *LockArr[T]) Get() []T {
	la.mu.Lock()
	defer la.mu.Unlock()
	return la.arr
}

func Last[T any](ts []T) T {
	if len(ts) == 0 {
		return *new(T)
	}
	return ts[len(ts)-1]
}
