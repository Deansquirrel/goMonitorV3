package taskService

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/global"
	"github.com/Deansquirrel/goMonitorV3/object"
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/Deansquirrel/goMonitorV3/worker"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/robfig/cron"
	"time"
)

type task struct {
	iTask ITask
	cType global.ConfigType
	hType global.HisType
}

func NewIntTask() *task {
	return newTask(&intTask{}, global.CInt, global.HInt)
}

func newTask(iTask ITask, cType global.ConfigType, hType global.HisType) *task {
	return &task{
		iTask: iTask,
		cType: cType,
		hType: hType,
	}
}

func (t *task) StartTask() error {
	rep := repository.NewIntConfigRepository()
	//获取配置列表
	list, err := rep.GetConfigList()
	if err != nil {
		return err
	}
	t.clearCache()

	t.startRegularRefresh()

	errMsg := ""
	errMsgFormat := "添加任务[%s]报错：%s；"
	for _, config := range list {
		err = t.addJob(config)
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, config.GetConfigId(), err.Error())
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (t *task) addJob(iConfig object.IConfigData) error {
	w, err := worker.NewWorker(iConfig)
	if err != nil {
		AddTask(iConfig.GetConfigId(), &TaskCache{
			Config:    iConfig,
			Cron:      nil,
			IsRunning: false,
		})
		return err
	}
	c := cron.New()
	AddTask(iConfig.GetConfigId(), &TaskCache{
		Config:    iConfig,
		Cron:      c,
		IsRunning: false,
	})
	err = c.AddJob(iConfig.GetSpec(), w)
	if err != nil {
		return nil
	} else {
		c.Start()
		UpdateTaskState(iConfig.GetConfigId(), true)
		return nil
	}
}

func (t *task) delJob(id string) {
	DelTask(id)
}

func (t *task) startJob(id string) error {
	list := GetTaskList()
	tc, ok := list[id]
	if ok {
		if tc.Cron == nil {
			return errors.New(fmt.Sprintf("配置无效，ID：%s", id))
		}
		if tc.IsRunning {
			return nil
		}
		tc.Cron.Start()
		UpdateTaskState(id, true)
		return nil
	} else {
		return errors.New(fmt.Sprintf("无效的ID：%s", id))
	}
}

func (t *task) stopJob(id string) error {
	list := GetTaskList()
	tc, ok := list[id]
	if ok {
		if tc.Cron == nil {
			return errors.New(fmt.Sprintf("配置无效，ID：%s", id))
		}
		if !tc.IsRunning {
			return nil
		}
		tc.Cron.Stop()
		UpdateTaskState(id, false)
		return nil
	} else {
		return errors.New(fmt.Sprintf("无效的ID：%s", id))
	}
}

//清除缓存配置
func (t *task) clearCache() {
	for _, id := range t.iTask.getCacheIdList() {
		t.delJob(id)
	}
}

func (t *task) startRegularRefresh() {
	c := cron.New()
	var err error
	err = c.AddFunc("0 0/1 * * * ?", t.refreshConfig)
	if err != nil {
		log.Error("添加配置刷新任务时遇到错误：" + err.Error())
	} else {
		log.Info("添加配置刷新任务完成")
	}
	err = c.AddFunc("0 0 0 * * ?", t.delHisData)
	if err != nil {
		log.Error("添加删除历史数据任务时遇到错误：" + err.Error())
	} else {
		log.Info("添加删除历史数据任务完成")
	}
	c.Start()
}

//刷新任务配置
func (t *task) refreshConfig() {
	err := t.refreshConfigWorker()
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
	}
}

//删除历史数据
func (t *task) delHisData() {
	d := time.Duration(1000 * 1000 * 1000 * 60 * 60 * 24 * global.SysConfig.TaskConfig.KeepDays)
	var rep repository.IHisRepository
	switch t.hType {
	case global.HInt:
		rep = repository.NewIntHisRepository()
	default:
		log.Error("未预知的hType")
		return
	}
	err := rep.ClearHis(d)
	if err != nil {
		log.Error("删除历史数据时遇到错误：" + err.Error())
	}
}

func (t *task) refreshConfigWorker() error {
	var rep repository.IConfigRepository
	switch t.cType {
	case global.CInt:
		rep = repository.NewIntConfigRepository()
	default:
		errMsg := "未预知的cType"
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	//获取配置列表
	list, err := rep.GetConfigList()
	if err != nil {
		return err
	}
	idList := make([]string, 0)
	idMap := make(map[string]object.IConfigData, 0)
	for _, iConfig := range list {
		config, ok := iConfig.(object.IConfigData)
		if !ok {
			return errors.New("不是有效的IConfigData")
		}
		idList = append(idList, config.GetConfigId())
		idMap[config.GetConfigId()] = config
	}

	addList, delList, checkList := goToolCommon.CheckDiff(idList, t.iTask.getCacheIdList())

	errMsg := ""
	errMsgFormat := "刷新任务[%s]报错：%s；"

	for _, id := range addList {
		err = t.addJob(idMap[id])
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
		}
	}
	for _, id := range delList {
		t.delJob(id)
	}
	for _, id := range checkList {
		tcList := GetTaskList()
		tc := tcList[id]
		if !tc.Config.IsEqual(idMap[id]) {
			t.delJob(id)
			err = t.addJob(idMap[id])
			if err != nil {
				errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
			}
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
