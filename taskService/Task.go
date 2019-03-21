package taskService

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
)

type task struct {
	config repository.IConfig
	t      ITask
}

func NewTask(t ITask, config repository.IConfig) *task {
	return &task{
		config: config,
		t:      t,
	}
}

func (t *task) getTaskConfigId(config repository.IConfigData) (string, error) {
	var v interface{}
	v = config
	switch c := v.(type) {
	case repository.IntConfigData:
		return c.FId, nil
	default:
		log.Warn("未预知的类型")
		return "", errors.New("未预知的类型")
	}
}

func (t *task) StartTask() error {
	rep := repository.NewConfigRepository(&repository.IntConfig{})
	//获取配置列表
	list, err := rep.GetConfigList()
	if err != nil {
		return err
	}
	//缓存配置列表、任务列表
	errMsg := ""
	errMsgFormat := "添加任务报错：%s，ID：%s；"
	for _, config := range list {
		err := t.addTask(config)
		if err != nil {
			id, _ := t.getTaskConfigId(config)
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, err.Error(), id)
		}
	}
	if errMsg != "" {
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	t.startRegularRefresh()
	return nil
}

func (t *task) startRegularRefresh() {

}

func (t *task) addTask(config repository.IConfigData) error {
	//------------------------------------------------------------------------------------------------------------------
	configStr, err := goToolCommon.GetJsonStr(config)
	if err != nil {
		log.Warn(fmt.Sprintf("Add Task，转换配置内容时遇到错误:%s", configStr))
		return err
	} else {
		log.Warn(fmt.Sprintf("Add Task:%s", configStr))
	}
	//------------------------------------------------------------------------------------------------------------------
	return nil
}
