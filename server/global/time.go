package global

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type MySQLTime struct {
	time.Time
}

func NewMySQLTime(t time.Time) *MySQLTime {
	return &MySQLTime{
		Time: t,
	}
}

func Now() *MySQLTime {
	return &MySQLTime{
		Time: time.Now(),
	}
}

func NowMySQLTime() MySQLTime {
	return MySQLTime{
		Time: time.Now(),
	}
}

func (t *MySQLTime) GetTime() time.Time {
	return t.Time
}

func NowMySQLTimeAdd(duration time.Duration) *MySQLTime {
	return &MySQLTime{
		Time: time.Now().Add(duration),
	}
}

func (t *MySQLTime) FormatTime() string {
	if t == nil || t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02 15:04:05")
}

func (t *MySQLTime) FormatDate() string {
	if t == nil || t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func (t *MySQLTime) AddDays(days int) *MySQLTime {
	if t == nil {
		return t
	}
	return &MySQLTime{
		Time: t.Time.AddDate(0, 0, days),
	}
}

// JSON 反序列化
func (t *MySQLTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		t.Time = time.Time{}
		return nil
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// JSON 序列化
func (t MySQLTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(`null`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Time.Format("2006-01-02 15:04:05"))), nil
}

// GORM 保存到数据库
func (t MySQLTime) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time.Format("2006-01-02 15:04:05"), nil
}

// GORM 从数据库读取
func (t *MySQLTime) Scan(v interface{}) error {
	switch val := v.(type) {
	case time.Time:
		t.Time = val
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		t.Time = tt
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		t.Time = tt
	default:
		return fmt.Errorf("cannot scan type %T into MySQLTime", v)
	}
	return nil
}
