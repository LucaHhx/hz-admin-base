package utils

import "github.com/shopspring/decimal"

type Number interface {
	Int | float32 | float64
}
type Signed interface {
	int8 | int32 | int | int64 | float64
}

func Sum[T Number](num ...T) float64 {
	if len(num) == 0 {
		return 0
	}
	var sum decimal.Decimal
	for _, v := range num {
		sum = sum.Add(decimal.NewFromFloat(float64(v)))
	}
	f, _ := sum.Float64()
	return f
}

func Mul[T Number, V Number](d1 T, num ...V) float64 {
	if len(num) == 0 || d1 == 0 {
		return 0
	}
	var sum = decimal.NewFromFloat(float64(d1))
	for _, v := range num {
		if v == 0 {
			return 0
		}
		sum = sum.Mul(decimal.NewFromFloat(float64(v)))
	}
	f, _ := sum.Float64()
	return f
}

func Div[T Number, V Number](d1 T, num ...V) float64 {
	if len(num) == 0 || d1 == 0 {
		return 0
	}
	var sum = decimal.NewFromFloat(float64(d1))
	for _, v := range num {
		if v == 0 {
			return 0
		}
		sum = sum.Div(decimal.NewFromFloat(float64(v)))
	}
	f, _ := sum.Float64()
	return f
}

func DivToInt[T Number, V Number](d1 T, num ...V) int64 {
	if len(num) == 0 || d1 == 0 {
		return 0
	}
	var sum = decimal.NewFromFloat(float64(d1))
	for _, v := range num {
		if v == 0 {
			return 0
		}
		sum = sum.Div(decimal.NewFromFloat(float64(v)))
	}
	return sum.IntPart()
}

func MulToInt[T Number, V Number](d1 T, num ...V) int64 {
	if len(num) == 0 || d1 == 0 {
		return 0
	}
	var sum = decimal.NewFromFloat(float64(d1))
	for _, v := range num {
		if v == 0 {
			return 0
		}
		sum = sum.Mul(decimal.NewFromFloat(float64(v)))
	}
	return sum.IntPart()
}

func Abs[T Signed](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

// NearKey 匹配最近值的key nums为降序
func NearKey[T Signed](nums []T, num T) []int {
	minDiff := nums[0] - num
	keys := []int{0}
	for i := 1; i < len(nums); i++ {
		diff := Abs(nums[i] - num)
		if diff < minDiff {
			minDiff = diff
			keys = []int{i}
		} else if diff == minDiff {
			keys = append(keys, i)
		} else {
			break
		}
	}
	return keys
}

// NearVal 匹配最近值 nums为降序
func NearVal[T Signed](nums []T, num T) T {
	minDiff := nums[0] - num
	for i := 1; i < len(nums); i++ {
		diff := Abs(nums[i] - num)
		if diff > minDiff {
			return nums[i-1]
		} else {
			minDiff = diff
		}
	}
	return nums[len(nums)-1]
}

func Mul100[T Number](num T) int64 {
	return decimal.NewFromFloat(float64(num)).Mul(decimal.NewFromInt(100)).IntPart()
}

func Div100[T Number](num T) float64 {
	v, _ := decimal.NewFromFloat(float64(num)).Div(decimal.NewFromInt(100)).Float64()
	return v
}

func Div100ToInt[T Number](num T) int64 {
	v := decimal.NewFromFloat(float64(num)).Div(decimal.NewFromInt(100)).IntPart()
	return v
}

func FlotDiv100(num float64) float64 {
	v, _ := decimal.NewFromFloat(num).Div(decimal.NewFromInt(100)).Float64()
	return v
}

func Avg[T Number](num ...T) float64 {
	if len(num) == 0 {
		return 0
	}
	var sum decimal.Decimal
	for _, v := range num {
		sum = sum.Add(decimal.NewFromFloat(float64(v)))
	}
	avg, _ := sum.Div(decimal.NewFromInt(int64(len(num)))).Float64()
	return avg
}

func Range(n1, n2 int) []int {
	numbers := make([]int, n2-n1+1)

	for i := range numbers {
		numbers[i] = n1 + i
	}
	return numbers
}

// DivRound 除后并保留小数点后n位
func DivRound[T Number, V Number](d1 T, num V, round int32) float64 {
	if d1 == 0 || num == 0 {
		return 0
	}
	f, _ := decimal.NewFromFloat(float64(d1)).
		Div(decimal.NewFromFloat(float64(num))).
		Round(round).Float64()
	return f
}

func IsNumber(v interface{}) (float64, bool) {
	switch value := v.(type) {
	case int:
		return float64(value), true
	case int8:
		return float64(value), true
	case int16:
		return float64(value), true
	case int32:
		return float64(value), true
	case int64:
		return float64(value), true
	case uint:
		return float64(value), true
	case uint8:
		return float64(value), true
	case uint16:
		return float64(value), true
	case uint32:
		return float64(value), true
	case uint64:
		return float64(value), true
	case float32:
		return float64(value), true
	case float64:
		return value, true
	default:
		return 0, false
	}
}

func RtxFloat[T Number](num T) float64 {

	//f, _ := decimal.NewFromFloat(float64(num)).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(6))).Round(6).Float64()
	return PowFloat(num, 6)
}

func PowFloat[T Number, N Int](num T, n N) float64 {
	f, _ := decimal.NewFromFloat(float64(num)).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(n)))).Round(6).Float64()
	return f
}

func DivToIntRate(f1, f2 float64) int64 {
	if f1 == 0 || f2 == 0 {
		return 0
	}
	return decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Mul(decimal.NewFromInt(100)).IntPart()
}
