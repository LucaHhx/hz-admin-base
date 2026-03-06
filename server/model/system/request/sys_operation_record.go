package request

import (
	"hz-admin-base/model/common/request"
	"hz-admin-base/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
