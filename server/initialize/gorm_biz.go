package initialize

import (
	"hab/global"
	"hab/model/tests"
)

func bizModel() error {
	db := global.HAB_DB
	err := db.AutoMigrate(tests.Order{})
	if err != nil {
		return err
	}
	return nil
}
