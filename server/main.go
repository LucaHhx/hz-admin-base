package main

import (
	"hz-admin-base/core"
	"hz-admin-base/global"
	"hz-admin-base/initialize"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	core.InitServer()

	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 设置时区
	var err error
	if time.Local, err = time.LoadLocation(global.GVA_CONFIG.System.TimeZone); err != nil {
		panic(err)
	}

	global.GVA_LOG.Info("nowTime: " + time.Now().Format("2006-01-02 15:04:05"))

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
