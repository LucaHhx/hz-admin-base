package utils

import (
	"hz-admin-base/global"
	"github.com/samber/lo"
	"reflect"
	"strings"
)

type ParamsComments struct {
	Name    string           `json:"name"`
	Comment string           `json:"comment"`
	Type    string           `json:"type,omitempty"`
	Sub     []ParamsComments `json:"sub,omitempty"`
}

type ParStruct struct {
	Name    string           `json:"name"`
	Comment string           `json:"comment"`
	Field   []ParamsComments `json:"field,omitempty"`
	Sub     []ParamsComments `json:"sub,omitempty"`
}

// 从 gorm 标签中提取 comment
func getCommentFromGormTag(gormTag string) string {
	parts := strings.Split(gormTag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "comment:") {
			return strings.TrimPrefix(part, "comment:")
		}
	}
	return ""
}

func NewParStruct(v interface{}, note string) *ParStruct {
	par := ParStruct{
		Name:    reflect.TypeOf(v).Name(),
		Comment: note,
	}
	pars := AnalysisStructure(v, "")
	par.Field = lo.Filter(pars, func(item ParamsComments, index int) bool {
		return len(item.Sub) == 0
	})

	par.Sub = lo.Filter(pars, func(item ParamsComments, index int) bool {
		return len(item.Sub) > 0
	})
	return &par
}

// AnalysisStructure 递归提取结构体或结构体指针的所有字段，包括嵌套字段。
func AnalysisStructure(v interface{}, prefix string) []ParamsComments {
	var list []ParamsComments = make([]ParamsComments, 0)
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		fName := field.Name
		gormTag := field.Tag.Get("gorm")
		if prefix != "" {
			fName = prefix + "." + field.Name
		}
		fName = strings.ReplaceAll(fName, "GVA_MODEL.", "")
		pre := ParamsComments{
			Name:    fName,
			Type:    fieldValue.Type().String(),
			Comment: getCommentFromGormTag(gormTag),
			Sub:     make([]ParamsComments, 0),
		}
		if fieldValue.Type() == reflect.TypeOf(global.GVA_MODEL{}) ||
			fieldValue.Type() == reflect.TypeOf(&global.GVA_MODEL{}).Elem() {
			pre = ParamsComments{
				Name:    strings.ReplaceAll(fName+".ID", "GVA_MODEL.", ""),
				Type:    "uint",
				Comment: "ID",
			}
		} else if fieldValue.Kind() == reflect.Struct {
			pre.Sub = append(pre.Sub, AnalysisStructure(fieldValue.Interface(), fName)...)
		} else if fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Struct {
			pre.Sub = append(pre.Sub, AnalysisStructure(fieldValue.Interface(), fName)...)
		}
		list = append(list, pre)
	}
	return list
}
