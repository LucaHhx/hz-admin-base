package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func SqlReplace(query string, vars []interface{}) string {
	var sb strings.Builder
	i := 0
	for _, ch := range query {
		if ch == '?' && i < len(vars) {
			sb.WriteString(quoteValue(vars[i]))
			i++
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func quoteValue(v interface{}) string {
	if v == nil {
		return "NULL"
	}

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return "NULL"
	}

	// 解引用指针
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return "NULL"
		}
		rv = rv.Elem()
		v = rv.Interface()
	}

	switch val := v.(type) {
	case time.Time:
		return fmt.Sprintf("'%s'", val.Format("2006-01-02 15:04:05"))
	case string:
		return fmt.Sprintf("'%s'", strings.ReplaceAll(val, "'", "''"))
	case []byte:
		return fmt.Sprintf("'%s'", strings.ReplaceAll(string(val), "'", "''"))
	case bool:
		if val {
			return "TRUE"
		}
		return "FALSE"
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", val)

	// sql.Null 系列类型
	case sql.NullString:
		if val.Valid {
			return fmt.Sprintf("'%s'", strings.ReplaceAll(val.String, "'", "''"))
		}
		return "NULL"
	case sql.NullInt64:
		if val.Valid {
			return fmt.Sprintf("%d", val.Int64)
		}
		return "NULL"
	case sql.NullFloat64:
		if val.Valid {
			return fmt.Sprintf("%f", val.Float64)
		}
		return "NULL"
	case sql.NullBool:
		if val.Valid {
			if val.Bool {
				return "TRUE"
			}
			return "FALSE"
		}
		return "NULL"
	case sql.NullTime:
		if val.Valid {
			return fmt.Sprintf("'%s'", val.Time.Format("2006-01-02 15:04:05"))
		}
		return "NULL"

	default:
		// 其他复杂类型尝试转为 JSON（如 map、struct、slice）
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("'%v'", v)
		}
		return fmt.Sprintf("'%s'", strings.ReplaceAll(string(b), "'", "''"))
	}
}
