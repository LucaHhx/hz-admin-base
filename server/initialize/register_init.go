package initialize

import (
	_ "hz-admin-base/source/example"
	_ "hz-admin-base/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
