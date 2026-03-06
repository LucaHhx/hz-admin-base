package system

import (
	"hz-admin-base/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
