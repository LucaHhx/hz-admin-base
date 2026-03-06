package v1

import (
	"hab/api/v1/api"
	"hab/api/v1/business"
	"hab/api/v1/system"
	"hab/api/v1/tests"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	BusinessApiGroup business.ApiGroup
	ApiApiGroup      api.ApiGroup
	TestsApiGroup    tests.ApiGroup
}
