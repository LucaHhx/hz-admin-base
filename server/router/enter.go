package router

import (
	"hab/router/api"
	"hab/router/business"
	"hab/router/system"
	"hab/router/tests"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   system.RouterGroup
	Business business.RouterGroup
	Api      api.RouterGroup
	Tests    tests.RouterGroup
}
