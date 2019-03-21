package global

import (
	"context"
	"github.com/Deansquirrel/goMonitorV3/config"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/robfig/cron"
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

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()

type TaskCache struct {
	Config    interface{}
	Cron      *cron.Cron
	IsRunning bool
}

var cacheList map[string]*TaskCache

func AddTask(id string, cache *TaskCache) {
	cacheList[id] = cache
}

func DelTask(id string) {
	_, ok := cacheList[id]
	if ok {
		delete(cacheList, id)
	}
}

func UpdateTaskState(id string, isRunning bool) {
	cache, ok := cacheList[id]
	if ok {
		cache.IsRunning = isRunning
	}
}

func init() {
	goToolMSSql.SetMaxIdleConn(15)
	goToolMSSql.SetMaxOpenConn(15)
	goToolMSSql.SetMaxLifetime(time.Second * 60)

	cacheList = make(map[string]*TaskCache)
}
