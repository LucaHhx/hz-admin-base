package request

import (
	"hz-admin-base/model/common/request"
	"hz-admin-base/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
