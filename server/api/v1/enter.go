package v1

import (
	"hz-admin-base/api/v1/api"
	"hz-admin-base/api/v1/business"
	"hz-admin-base/api/v1/example"
	"hz-admin-base/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	BusinessApiGroup business.ApiGroup
	ApiApiGroup      api.ApiGroup
}
