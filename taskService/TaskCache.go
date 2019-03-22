package taskService

import (
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/robfig/cron"
	"sync"
)

var lock sync.Mutex

var cacheList map[string]*TaskCache

func init() {
	cacheList = make(map[string]*TaskCache)
}

type TaskCache struct {
	Config    repository.IConfigData
	Cron      *cron.Cron
	IsRunning bool
}

func GetTaskList() map[string]*TaskCache {
	lock.Lock()
	defer lock.Unlock()
	return cacheList
}

func AddTask(id string, cache *TaskCache) {
	lock.Lock()
	defer lock.Unlock()
	cacheList[id] = cache
}

func DelTask(id string) {
	lock.Lock()
	defer lock.Unlock()
	t, ok := cacheList[id]
	if ok {
		t.IsRunning = false
		t.Cron.Stop()
		delete(cacheList, id)
	}
}

func UpdateTaskState(id string, isRunning bool) {
	lock.Lock()
	defer lock.Unlock()
	cache, ok := cacheList[id]
	if ok {
		cache.IsRunning = isRunning
	}
}
