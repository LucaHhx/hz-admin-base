package request

import (
	"hab/model/common/request"
	"hab/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
