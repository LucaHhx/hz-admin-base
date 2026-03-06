package convert

import (
	"reflect"
	"strings"
	"sync"
)

type fieldInfo struct {
	Index    []int  // 字段索引路径
	Key      string // 映射的键名
	TagValue string // 标签值
}

var (
	// 字段缓存
	fieldCache     sync.Map
	defaultTagName = "json"
)

// StructToMap 将结构体转换为map，支持标签指定
func StructToMap(obj any, tag ...string) (m map[string]any) {
	m = make(map[string]any)
	if obj == nil {
		return m
	}

	// 解析标签
	tagName := ""
	if len(tag) == 0 {
		tagName = defaultTagName
	} else {
		tagName = tag[0]
	}

	val := reflect.ValueOf(obj)
	// 解引用指针
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return m
		}
		val = val.Elem()
	}

	// 仅处理结构体
	if val.Kind() != reflect.Struct {
		return m
	}

	// 获取字段信息
	infos := getFieldInfos(val.Type(), tagName)

	// 构建结果map
	for _, info := range infos {
		// 处理"-"忽略标记
		if info.TagValue == "-" {
			continue
		}

		fieldVal := val.FieldByIndex(info.Index)
		// 添加字段到map
		m[info.Key] = fieldVal.Interface()
	}

	return m
}

// 获取结构体的字段信息（带缓存）
func getFieldInfos(typ reflect.Type, tagName string) []fieldInfo {
	// 生成缓存键
	cacheKey := struct {
		Type    reflect.Type
		TagName string
	}{typ, tagName}

	// 尝试从缓存读取
	if v, ok := fieldCache.Load(cacheKey); ok {
		return v.([]fieldInfo)
	}

	// 缓存不存在时解析字段
	fields := make([]fieldInfo, 0, typ.NumField())
	queue := []reflect.Type{typ}

	// 使用队列处理嵌套结构体
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < current.NumField(); i++ {
			field := current.Field(i)

			// 忽略未导出字段
			if field.PkgPath != "" {
				continue
			}

			// 处理结构体标签
			tagStr := field.Tag.Get(tagName)
			if tagStr == "" && tagName != defaultTagName {
				tagStr = field.Tag.Get(defaultTagName)
			}

			// 处理匿名字段（展开嵌套结构体）
			if field.Anonymous {
				ft := field.Type
				// 解引用指针
				if ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}
				// 只处理结构体类型的匿名字段
				if ft.Kind() == reflect.Struct {
					queue = append(queue, ft)
					continue
				}
			}

			// 解析标签选项
			key, opts := parseTag(tagStr)
			if key == "" {
				key = field.Name
			}

			// 创建字段信息
			info := fieldInfo{
				Index:    field.Index,
				Key:      key,
				TagValue: opts,
			}

			fields = append(fields, info)
		}
	}

	// 缓存解析结果
	fieldCache.Store(cacheKey, fields)
	return fields
}

// 解析结构体标签（简化版）
func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}
