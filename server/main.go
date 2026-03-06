package main

import (
	"hab/core"
	"hab/global"
	"hab/initialize"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	core.InitServer()

	go func() {
		global.HAB_LOG.Info("Starting pprof server on :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			global.HAB_LOG.Error("pprof server failed", zap.Error(err))
		}
	}()

	// 设置时区
	var err error
	if time.Local, err = time.LoadLocation(global.HAB_CONFIG.System.TimeZone); err != nil {
		panic(err)
	}

	global.HAB_LOG.Info("nowTime: " + time.Now().Format("2006-01-02 15:04:05"))

	// 启动后台服务（阻塞）
	switch global.SysMode {
	case global.ModeAll:
		go RunApi()
		RunBackend()
	case global.ModeBackend:
		RunBackend()
	case global.ModeApi:
		RunApi()
	}
}

func RunBackend() {
	initialize.Timer() // 初始化定时任务
	core.RunWindowsServer()
}

func RunApi() {
	core.RunApiServer()
}
