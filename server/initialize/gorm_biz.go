package initialize

import (
	"hab/global"
)

func bizModel() error {
	db := global.HAB_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
