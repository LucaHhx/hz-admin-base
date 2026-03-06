package utils

import (
	"errors"
	"fmt"
	"hz-admin-base/global"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Field struct {
	Name    string
	Value   reflect.Value
	Comment string
	IsArr   bool
}

type Expression struct {
	value      any
	analysis   map[string]Field
	Expression string
}

func NewExpression(value any) (*Expression, error) {
	exp := &Expression{
		value:    value,
		analysis: make(map[string]Field),
	}
	err := exp.AnalysisStructure(value, "")
	if err != nil {
		global.GVA_LOG.Error("AnalysisStructure error", zap.Error(err))
		return nil, err
	}
	return exp, nil
}

func (e *Expression) GetAnalysis() map[string]Field {
	return e.analysis
}

// AnalysisStructure 递归提取结构体或结构体指针的所有字段，包括嵌套字段。
func (e *Expression) AnalysisStructure(v interface{}, prefix string) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return errors.New("value is not a struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Construct the full field name with prefix
		fName := field.Name
		if prefix != "" {
			fName = prefix + "." + field.Name
		}
		fName = strings.ReplaceAll(fName, "GVA_MODEL.", "")
		gormTag := field.Tag.Get("gorm")
		// Check if the field is a struct or a pointer to a struct
		if fieldValue.Type() == reflect.TypeOf(time.Time{}) || fieldValue.Type() == reflect.TypeOf(&time.Time{}) {
			continue
		}

		if fieldValue.Kind() == reflect.Struct {
			// If it's a struct, recursively call AnalysisStructure
			err := e.AnalysisStructure(fieldValue.Interface(), fName)
			if err != nil {
				return err
			}
		} else if fieldValue.Kind() == reflect.Slice {
			e.analysis[fName] = Field{
				Name:    fName,
				Value:   fieldValue,
				Comment: getCommentFromGormTag(gormTag),
			}
			for j := 0; j < fieldValue.Len(); j++ {
				e.analysis[fmt.Sprintf("%s[%d]", fName, j)] = Field{
					Name:    fmt.Sprintf("%s[%d]", fName, j),
					Value:   fieldValue.Index(j),
					Comment: getCommentFromGormTag(gormTag),
					IsArr:   true,
				}
			}
		} else if fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Struct {
			// If it's a pointer to a struct, recursively call AnalysisStructure on the dereferenced value
			err := e.AnalysisStructure(fieldValue.Interface(), fName)
			if err != nil {
				return err
			}
		} else {
			// Otherwise, add the field directly to the map
			e.analysis[fName] = Field{
				Name:    fName,
				Value:   fieldValue,
				Comment: getCommentFromGormTag(gormTag),
			}
		}
	}
	return nil
}

func (e *Expression) Evaluate(expression string) (bool, error) {
	exp, err := e.ParseContainsExpression(expression)
	if err != nil {
		global.GVA_LOG.Error("ParseContainsExpression error", zap.Error(err), zap.String("expression", expression))
		return false, err
	}
	e.Expression = e.ReplaceFieldsWithValues(exp)
	// Use Go's parser and evaluator to determine if the expression is true
	fset := token.NewFileSet()
	_, err = parser.ParseExpr(e.Expression)
	if err != nil {
		global.GVA_LOG.Error("failed to parse expression", zap.Error(err), zap.String("expression", e.Expression))
		return false, err
	}

	// Evaluate the parsed expression
	result, err := types.Eval(fset, nil, token.NoPos, e.Expression)
	if err != nil {
		global.GVA_LOG.Error("failed to evaluate expression", zap.Error(err), zap.String("expression", e.Expression))
		return false, err
	}

	// Convert the result to a boolean value
	value, err := strconv.ParseBool(result.Value.String())
	if err != nil {
		global.GVA_LOG.Error("failed to convert result to boolean", zap.Error(err), zap.String("expression", e.Expression))
		return false, err
	}
	global.GVA_LOG.Info("expression evaluation result", zap.String("expression", e.Expression), zap.Bool("result", value))
	return value, nil
}

func (e *Expression) ParseContainsExpression(expression string) (string, error) {
	// D定义包含函数的正则表达式，例如：, contains(Tags, 'test')
	re := regexp.MustCompile(`contains\(([^,]+),\s*"([^']+)"\)`)
	matches := re.FindAllStringSubmatch(expression, -1)

	// Replace each contains match with equivalent boolean OR expression
	for _, match := range matches {
		containsExpression := make([]string, 0)
		field := match[1]
		value := match[2]

		field = strings.ReplaceAll(strings.ReplaceAll(field, "{", ""), "}", "")
		tValue, ok := e.analysis[field]
		if !ok {
			return "", fmt.Errorf("field %s not found in analysis", field)
		}

		for i := range tValue.Value.Len() {
			containsExpression = append(containsExpression, fmt.Sprintf("{%s[%d]} == \"%s\"", field, i, value))
		}
		var cExpStr string
		if len(containsExpression) > 0 {
			cExpStr = fmt.Sprintf("(%s)", ArrJoin(containsExpression, "||"))
		} else {
			cExpStr = fmt.Sprintf("%s", "false")
		}
		expression = strings.Replace(expression, match[0], cExpStr, 1)
	}
	return expression, nil
}

func (e *Expression) ReplaceFieldsWithValues(expression string) string {
	for s, value := range e.analysis {
		if strings.Contains(expression, fmt.Sprintf("{%s}", s)) {
			expression = strings.ReplaceAll(expression, fmt.Sprintf("{%s}", s), GetKindValue(value.Value))
		}
	}
	return expression
}

func GetKindValue(fieldValue reflect.Value) string {
	switch fieldValue.Kind() {
	case reflect.Float64, reflect.Float32:
		if fieldValue.CanFloat() {
			return fmt.Sprintf("%.2f", fieldValue.Float())
		}
		return fmt.Sprintf("%f", 0.0)
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if fieldValue.CanInt() {
			return fmt.Sprintf("%d", fieldValue.Int())
		}
		return fmt.Sprintf("%d", 0)
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		if fieldValue.CanUint() {
			return fmt.Sprintf("%d", fieldValue.Uint())
		}
		return fmt.Sprintf("%d", 0)

	case reflect.String:
		if fieldValue.CanInterface() {
			return fmt.Sprintf("\"%s\"", fieldValue.String())
		}
		return ""
	case reflect.Bool:
		return fmt.Sprintf("%t", fieldValue.Bool())
	default:
		return "NaN"
	}
}
