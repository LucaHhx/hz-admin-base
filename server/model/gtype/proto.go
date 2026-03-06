package gtype

import (
	"database/sql/driver"
	"google.golang.org/protobuf/proto"
	"hab/enum"
)

type Proto[T proto.Message] struct {
	Data T `json:"data" form:"data"` // Data is the main content of the proto message
}

func (a *Proto[T]) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return enum.Msg_InvalidParams
	}
	newA := new(Proto[T])
	*a = *newA
	if string(data) == "null" {
		return nil
	}
	proto.Unmarshal(data, a.Data)
	return nil
}

func (a Proto[T]) Value() (driver.Value, error) {
	return proto.Marshal(a.Data)
}

func (a Proto[T]) Info() T {
	return a.Data
}
