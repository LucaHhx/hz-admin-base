package timedtask

import (
	"hab/global"

	"go.uber.org/zap"
)

func BackendTimerTask() {
	// 测试定时任务 每1分钟一次
	if _, err := global.HAB_Timer.AddTaskByFunc("test", "* * * * *", test, "test"); err != nil {
		global.HAB_LOG.Error("add timer test error:", zap.Error(err))
	}
}

func test() {
	//global.HAB_LOG.Debug("test")
}
