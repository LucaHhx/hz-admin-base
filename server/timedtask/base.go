package timedtask

import (
	"hz-admin-base/global"

	"go.uber.org/zap"
)

func BackendTimerTask() {
	// 测试定时任务 每1分钟一次
	if _, err := global.GVA_Timer.AddTaskByFunc("test", "* * * * *", test, "test"); err != nil {
		global.GVA_LOG.Error("add timer test error:", zap.Error(err))
	}
}

func test() {
	//global.GVA_LOG.Debug("test")
}
