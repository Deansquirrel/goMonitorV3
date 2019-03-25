package global

import (
	"context"
	"github.com/Deansquirrel/goMonitorV3/config"
	"github.com/Deansquirrel/goToolMSSql"
	"time"
)

const (
	//PreVersion = "0.0.3 Build20190315"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"
)

const (
	HttpConnectTimeout = 30
)

type ConfigType int

const (
	CMConfig ConfigType = iota
	CInt
	CCrmDzXf
	CHealth
	CWebState
)

type NotifyType int

const (
	NotifyDingTalkRobot NotifyType = iota
)

type WorkerType int

const (
	WMConfig WorkerType = iota
	WCrmDzXf
	WHealth
	WWebState
)

type HisType int

const (
	HInt HisType = iota
	HCrmDzXf
	HHealth
	HWebState
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()

func init() {
	goToolMSSql.SetMaxIdleConn(15)
	goToolMSSql.SetMaxOpenConn(15)
	goToolMSSql.SetMaxLifetime(time.Second * 60)
}
