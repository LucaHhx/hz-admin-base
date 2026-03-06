package utils

import (
	"encoding/json"
	"hz-admin-base/global"
	"strings"
)

func GetJSONKeys(jsonStr string) (keys []string, err error) {
	// 使用json.Decoder，以便在解析过程中记录键的顺序
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	// 确保数据是一个对象
	if t != json.Delim('{') {
		return nil, err
	}
	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			return nil, err
		}
		keys = append(keys, t.(string))

		// 解析值
		var value interface{}
		err = dec.Decode(&value)
		if err != nil {
			return nil, err
		}
	}
	return keys, nil
}

// JsonEncode 用sonic进行数据编码
func JsonEncode(data interface{}) ([]byte, error) {
	return global.Json.Marshal(&data)
}

// JsonDecode 用sonic进行数据解码
func JsonDecode(data []byte, to interface{}) error {
	return global.Json.Unmarshal(data, to)
}
