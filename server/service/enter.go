package service

import (
	"hab/service/business"
	"hab/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}
