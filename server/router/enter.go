package router

import (
	"hz-admin-base/router/api"
	"hz-admin-base/router/business"
	"hz-admin-base/router/example"
	"hz-admin-base/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Business business.RouterGroup
	Api      api.RouterGroup
}
