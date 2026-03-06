package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"hab/enum"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"

	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RegisterRule
//@description: 注册自定义规则方案建议在路由初始化层即注册
//@param: key string, rule Rules
//@return: err error

func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册,无法重复注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: NotEmpty
//@description: 非空 不能为其对应类型的0值
//@return: string

func NotEmpty() string {
	return "notEmpty"
}

// @author: [zooqkl](https://github.com/zooqkl)
// @function: RegexpMatch
// @description: 正则校验 校验输入项是否满足正则表达式
// @param:  rule string
// @return: string

func RegexpMatch(rule string) string {
	return "regexp=" + rule
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Lt
//@description: 小于入参(<) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Lt(mark string) string {
	return "lt=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Le
//@description: 小于等于入参(<=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Le(mark string) string {
	return "le=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Eq
//@description: 等于入参(==) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Eq(mark string) string {
	return "eq=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ne
//@description: 不等于入参(!=)  如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ne(mark string) string {
	return "ne=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ge
//@description: 大于等于入参(>=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ge(mark string) string {
	return "ge=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Gt
//@description: 大于入参(>) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Gt(mark string) string {
	return "gt=" + mark
}

//
//@author: [piexlmax](https://github.com/piexlmax)
//@function: Verify
//@description: 校验方法
//@param: st interface{}, roleMap Rules(入参实例，规则map)
//@return: err error

func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if tagVal.Type.Kind() == reflect.Struct {
			if err = Verify(val.Interface(), roleMap); err != nil {
				return err
			}
		}
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case strings.Split(v, "=")[0] == "regexp":
					if !regexpMatch(strings.Split(v, "=")[1], val.String()) {
						return errors.New(tagVal.Name + "格式校验不通过")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围," + v)
					}
				}
			}
		}
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compareVerify
//@description: 长度和数字的校验方法 根据类型自动校验
//@param: value reflect.Value, VerifyStr string
//@return: bool

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String:
		return compare(len([]rune(value.String())), VerifyStr)
	case reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: isBlank
//@description: 非空校验
//@param: value reflect.Value
//@return: bool

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compare
//@description: 比较函数
//@param: value interface{}, VerifyStr string
//@return: bool

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}

var valid = validator.New()

func BindAndValidate(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return enum.Msg_InvalidParams
	}
	if err := ValidateStruct(req); err != nil {
		return err
	}
	return nil
}

func BindAndValid(data []byte, req interface{}) error {
	if err := json.Unmarshal(data, req); err != nil {
		return enum.Msg_InvalidParams
	}
	if err := ValidateStruct(req); err != nil {
		return err
	}
	return nil
}

func BindMapValid(data map[string]interface{}, req interface{}) error {
	err := mapstructure.Decode(data, req)
	if err != nil {
		return enum.Msg_InvalidParams
	}
	if err := ValidateStruct(req); err != nil {
		return err
	}
	return nil
}

func ValidateStruct(s interface{}) error {
	err := valid.Struct(s)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return enum.NewCodeMsg(enum.Code_InvalidParams, err.Error())
		}
		for _, err := range err.(validator.ValidationErrors) {
			return enum.NewCodeMsg(enum.Code_InvalidParams, translateError(err))
		}
	}
	return nil
}

// translateError converts validation errors into user-friendly English messages
func translateError(fe validator.FieldError) string {
	// Use the label or field name
	fieldName := fe.Field()
	if label := fe.Tag(); label != "" {
		fieldName = fe.StructField() // Fallback to struct field name
	}

	// Map error tags to human-readable messages
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", fieldName)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long.", fieldName, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long.", fieldName, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address.", fieldName)
	case "oneof":
		return fmt.Sprintf("%s must be one of: (%s)", fieldName, fe.Param())
	default:
		return fmt.Sprintf("%s is invalid (%s).", fieldName, fe.Tag())
	}
}
