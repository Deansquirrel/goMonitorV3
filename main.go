package main

import (
	"context"
	"github.com/Deansquirrel/goMonitorV3/common"
	"github.com/Deansquirrel/goMonitorV3/configRepository"
	"github.com/Deansquirrel/goMonitorV3/global"
	"github.com/Deansquirrel/goToolCommon"
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
	//==================================================================================================================
	ic := configRepository.IntConfig{}
	r := configRepository.NewConfigRepository(&ic)
	list, err := r.GetConfigList()
	if err != nil {
		log.Error(err.Error())
		return
	}
	for _, val := range list {
		s, err := goToolCommon.GetJsonStr(val)
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Debug(s)
		}
	}
	//==================================================================================================================
}
