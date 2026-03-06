package gtype

import (
	"database/sql/driver"
	"encoding/json"
	"hz-admin-base/enum"
)

type Limit struct {
	Min int64 `json:"min" form:"min"`
	Max int64 `json:"max" form:"max"`
}

func (a *Limit) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return enum.Msg_InvalidParams
	}

	*a = Limit{}
	if string(data) == "null" {
		return nil
	}
	json.Unmarshal(data, a)
	return nil
}

func (a Limit) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a Limit) GormDataType() string {
	return "varchar(500)"
}

type Data[T any] map[string]any

func (a *Data[T]) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return enum.Msg_InvalidParams
	}

	*a = make(Data[T])
	if string(data) == "null" {
		return nil
	}
	json.Unmarshal(data, a)
	return nil
}

func (a Data[T]) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a Data[T]) Info() T {
	info := new(T)
	if len(a) == 0 {
		return *info
	}
	data, err := json.Marshal(&a)
	if err != nil {
		return *info
	}
	err = json.Unmarshal(data, info)
	if err != nil {
		return *info
	}
	return *info
}
