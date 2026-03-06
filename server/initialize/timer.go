package initialize

import (
	"hz-admin-base/timedtask"
)

func Timer() {
	// 后台定时任务
	timedtask.BackendTimerTask()
}
