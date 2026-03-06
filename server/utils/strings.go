package utils

import (
	"fmt"
	"hz-admin-base/enum"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/streadway/amqp"
)

type Int interface {
	int8 | int | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64
}

// String : Conver "val" to a String
func String(val interface{}) (string, error) {
	switch ret := val.(type) {
	case string:
		return ret, nil
	case []byte:
		return string(ret), nil
	case float64:
		return decimal.NewFromFloat(ret).String(), nil
	case int:
		return strconv.Itoa(ret), nil
	default:
		str := fmt.Sprintf("%+v", val)
		if val == nil || len(str) == 0 {
			return "", fmt.Errorf("conver.String(), the %+v is empty", val)
		}
		return str, nil
	}
}

// StringMust : Must Conver "val" to a String
func StringMust(val interface{}, def ...string) string {
	ret, err := String(val)
	if err != nil && len(def) > 0 {
		return def[0]
	}
	return ret
}

func StringToInt[T Int](num string) T {
	ret, _ := decimal.NewFromString(num)
	f, _ := ret.Float64()
	return T(f)
}

func IsEmptyStr(val *string) bool {
	return val == nil || len(*val) == 0
}

func ConvStr(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

// CutToInt 将字符串切成两个int
func CutToInt[T Int](s, sep string) (T, T, bool) {
	v1, v2, ok := strings.Cut(s, sep)
	if !ok {
		return 0, 0, false
	}
	return T(Atoi(v1)), T(Atoi(v2)), true
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Itoa[V Int](i V) string {
	return strconv.Itoa(int(i))
}

// RemoveBytes 移除字符串中的单个字符 性能较高
func RemoveBytes(s string, v byte) string {
	bytes := []byte(s)
	j := 0
	for _, b := range bytes {
		if b != v {
			bytes[j] = b
			j++
		}
	}
	return string(bytes[:j])
}

func StrReplace(s string, reps map[string]string) string {
	for k, v := range reps {
		s = strings.Replace(s, fmt.Sprintf("@%s", k), v, -1)
	}
	return s
}

func FormatCommand(s string) []string {
	// 所有制表符替换为单个空格
	reg := regexp.MustCompile("\\s+")
	s = reg.ReplaceAllString(s, " ")
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	return arr
}

func StringToFloat(s string) float64 {
	// 去除前后空格
	s = strings.TrimSpace(s)
	// 去除千分号 ","
	s = strings.ReplaceAll(s, ",", "")
	// 转换为 float64
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0 // 或者你可以选择 panic(err) 或返回 error 类型
	}
	return f
}

func RepeatCount(table amqp.Table) int64 {
	count, ok := table[enum.Head_RepeatCount].(int64)
	if !ok {
		return 0
	}
	return count
}

// FloatWithCommas 浮点添加逗号
func FloatWithCommas(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := strings.TrimRight(parts[1], "0")

	// 插入千分位
	n := len(intPart)
	if n <= 3 {
		if decPart == "" {
			return intPart
		}
		return intPart + "." + decPart
	}

	var b strings.Builder
	pre := n % 3
	if pre > 0 {
		b.WriteString(intPart[:pre])
		if n > pre {
			b.WriteString(",")
		}
	}

	for i := pre; i < n; i += 3 {
		b.WriteString(intPart[i : i+3])
		if i+3 < n {
			b.WriteString(",")
		}
	}

	if decPart == "" {
		return b.String()
	}
	return b.String() + "." + decPart
}

// FloatWithIndCommas 浮点转为印尼的带逗号格式
func FloatWithIndCommas(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := strings.TrimRight(parts[1], "0")

	// 插入千分位
	n := len(intPart)
	if n <= 3 {
		if decPart == "" {
			return intPart
		}
		return intPart + "," + decPart
	}

	var b strings.Builder
	pre := n % 3
	if pre > 0 {
		b.WriteString(intPart[:pre])
		if n > pre {
			b.WriteString(".")
		}
	}

	for i := pre; i < n; i += 3 {
		b.WriteString(intPart[i : i+3])
		if i+3 < n {
			b.WriteString(".")
		}
	}

	if decPart == "" {
		return b.String()
	}
	return b.String() + "," + decPart
}

// RemoveFloatZeros 去除浮点数末尾的零
func RemoveFloatZeros(s string) string {
	// 如果字符串包含小数点
	if strings.Contains(s, ".") {
		// 先去除末尾的零
		s = strings.TrimRight(s, "0")
		// 如果去除后末尾是小数点，也去除小数点
		s = strings.TrimRight(s, ".")
	}
	return s
}

// RemoveIndFloatZeros 去除印尼的带逗号格式的小数末尾的零
func RemoveIndFloatZeros(s string) string {
	// 如果字符串包含小数点
	if strings.Contains(s, ",") {
		// 先去除末尾的零
		s = strings.TrimRight(s, "0")
		// 如果去除后末尾是小数点，也去除小数点
		s = strings.TrimRight(s, ",")
	}
	return s
}

func IntWithCommas(n int) string {
	s := strconv.Itoa(n)
	nLen := len(s)
	if n < 0 {
		s = s[1:]
		nLen--
	}
	var b strings.Builder
	if n < 0 {
		b.WriteByte('-')
	}
	pre := nLen % 3
	if pre > 0 {
		b.WriteString(s[:pre])
		if nLen > pre {
			b.WriteString(",")
		}
	}
	for i := pre; i < nLen; i += 3 {
		b.WriteString(s[i : i+3])
		if i+3 < nLen {
			b.WriteString(",")
		}
	}
	return b.String()
}
