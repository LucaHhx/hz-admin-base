package system

import (
	"hab/global"
)

type JwtBlacklist struct {
	global.HAB_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
