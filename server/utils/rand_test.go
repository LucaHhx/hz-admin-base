package utils_test

import (
	"fmt"
	"hz-admin-base/utils"
	"testing"
)

func TestRandNum6(t *testing.T) {
	for i := range 35 {
		fmt.Println(fmt.Sprintf("第%d次生成的随机数为：%s", i, utils.RandNum6()))
	}

}
