package utils

import (
	"fmt"
	"hz-admin-base/global"
	"runtime"
	"strings"
)

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// PanicRecover 捕获恐慌
func PanicRecover() {
	if err := recover(); err != nil {
		global.GVA_LOG.Error("Analysis panic " + fmt.Sprintf("%v", err))
		global.GVA_LOG.Error(Stack())
		return
	}
}

func Stack() string {
	buf := make([]byte, 10000)
	n := runtime.Stack(buf, false)
	buf = buf[:n]

	s := string(buf)

	// skip nano frames lines
	const skip = 7
	count := 0
	index := strings.IndexFunc(s, func(c rune) bool {
		if c != '\n' {
			return false
		}
		count++
		return count == skip
	})
	return s[index+1:]
}

// ProcessInBatches 将数组按指定的 batchSize 进行分批处理
func ProcessInBatches[T any](arr []T, batchSize int, processor func([]T) error) (err error) {
	total := len(arr)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}
		// 执行批处理函数
		err = processor(arr[i:end])
		if err != nil {
			return err
		}
	}
	return nil
}
