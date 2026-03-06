package utils

import (
	"reflect"
	"sort"
	"strings"
)

func MakeOriginSign(params map[string]interface{}, secret string) string {
	var (
		sortedKeys []string
		s          string
		arr        []string
	)
	var newpParams = make(map[string]string)
	for k, v := range params {
		value := StringMust(v)
		if k == "sign" || value == "" {
			continue
		}
		newpParams[k] = value
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		arr = append(arr, k+"="+newpParams[k])
	}
	s = strings.Join(arr, "&") + secret
	//global.HAB_LOG.Info(s)
	return s
}

func Md5Sign(params map[string]interface{}, secret string) string {
	return MD5V([]byte(MakeOriginSign(params, secret)))
}

// StructToSpreadMap 接口体转Spread out的map
func StructToSpreadMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)

	// 如果它是指针，则取消引用
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		fieldValue := objValue.Field(i)

		// Check if the field is an embedded struct
		if fieldType.Anonymous && fieldValue.Kind() == reflect.Struct {
			// Recursively convert the nested struct to a map
			for k, v := range StructToSpreadMap(fieldValue.Interface()) {
				data[k] = v
			}
		} else {
			key := fieldType.Name
			if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
				key = jsonTag
			}
			data[key] = fieldValue.Interface()
		}
	}
	return data
}

// MakeSign 生成签名
func MakeSign(data any, secret string) string {
	return Md5Sign(StructToSpreadMap(data), secret)
}

func VerifySign(data any, sign string, secret string) bool {
	var lsing string
	switch data.(type) {
	case map[string]interface{}:
		lsing = Md5Sign(data.(map[string]interface{}), secret)
	default:
		lsing = MakeSign(data, secret)
	}
	return lsing == sign
}
