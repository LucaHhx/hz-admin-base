package initialize

import (
	"hab/global"
)

func bizModel() error {
	db := global.HAB_DB
	_ = db
	return nil
}
