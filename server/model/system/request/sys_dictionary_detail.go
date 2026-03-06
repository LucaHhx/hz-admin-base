package request

import (
	"hab/model/common/request"
	"hab/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
