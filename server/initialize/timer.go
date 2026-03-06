package initialize

import (
	"hab/timedtask"
)

func Timer() {
	// 后台定时任务
	timedtask.BackendTimerTask()
}
