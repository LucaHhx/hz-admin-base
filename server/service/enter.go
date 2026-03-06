package service

import (
	"hz-admin-base/service/business"
	"hz-admin-base/service/example"
	"hz-admin-base/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}
