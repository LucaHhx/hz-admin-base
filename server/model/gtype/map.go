package gtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Map[V comparable, T any] map[V]T

func (a *Map[V, T]) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	if len(data) == 0 {
		*a = make(Map[V, T])
		return nil
	}
	*a = make(Map[V, T])
	if string(data) == "null" {
		return nil
	}
	json.Unmarshal(data, a)
	return nil
}

func (a Map[V, T]) Value() (driver.Value, error) {
	return json.Marshal(a)
}
