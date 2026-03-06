package gtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Any[T any] struct {
	Data *T `json:"data" form:"data"`
}

func (a *Any[T]) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	*a = Any[T]{Data: new(T)}
	if len(data) == 0 {
		return nil
	}
	if string(data) == "null" {
		return nil
	}
	json.Unmarshal(data, a)
	return nil
}

func (a Any[T]) Value() (driver.Value, error) {
	value, err := json.Marshal(a)
	return value, err
}
