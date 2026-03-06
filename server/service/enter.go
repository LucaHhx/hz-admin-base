package service

import (
	"hab/service/business"
	"hab/service/system"
	"hab/service/tests"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	TestsServiceGroup    tests.ServiceGroup
}
