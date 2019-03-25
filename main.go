package main

import (
	"context"
	"github.com/Deansquirrel/goMonitorV3/common"
	"github.com/Deansquirrel/goMonitorV3/global"
	"github.com/Deansquirrel/goMonitorV3/taskService"
	log "github.com/Deansquirrel/goToolLog"
)

func main() {
	//==================================================================================================================
	log.Warn("程序启动")
	defer log.Warn("程序退出")
	//==================================================================================================================
	config, err := common.GetSysConfig("config.toml")
	if err != nil {
		log.Error("加载配置文件时遇到错误：" + err.Error())
		return
	}
	config.FormatConfig()
	global.SysConfig = config
	err = common.RefreshSysConfig(*global.SysConfig)
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
		return
	}
	global.Ctx, global.Cancel = context.WithCancel(context.Background())

	err = taskService.NewIntTask().StartTask()
	if err != nil {
		log.Debug(err.Error())
	}

	select {
	case <-global.Ctx.Done():
	}

	//==================================================================================================================
}
