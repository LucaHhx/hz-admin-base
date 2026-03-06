package gtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Arr[T any] []T

func (a *Arr[T]) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	*a = make(Arr[T], 0)
	if string(data) == "null" {
		return nil
	}
	json.Unmarshal(data, a)
	if a == nil {
		*a = make(Arr[T], 0)
	}
	return nil
}

func (a Arr[T]) Value() (driver.Value, error) {
	return json.Marshal(a)
}
