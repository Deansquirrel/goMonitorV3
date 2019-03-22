package worker

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/notify"
	"github.com/Deansquirrel/goMonitorV3/repository"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
)

type worker struct {
	iConfig repository.IConfigData
	iWorker IWorker
}

func NewWorker(iConfig repository.IConfigData) (*worker, error) {
	workerRunner, err := getWorker(iConfig)
	if err != nil {
		return nil, err
	}
	return &worker{
		iConfig: iConfig,
		iWorker: workerRunner,
	}, nil
}

func getWorker(iConfig repository.IConfigData) (IWorker, error) {
	switch reflect.TypeOf(iConfig).String() {
	case "*repository.IntConfigData":
		config, ok := iConfig.(*repository.IntConfigData)
		if ok {
			return newIntWorker(config), nil
		} else {
			return nil, errors.New("强制类型转换失败[IntConfigData]")
		}
	default:
		return nil, errors.New("未预知的配置类型")
	}
}

func (w *worker) Run() {
	msg, hisData := w.iWorker.GetMsg()
	defer func() {
		if hisData != nil {
			err := w.iWorker.SaveSearchResult(hisData)
			if err != nil {
				w.sendMsg(w.iConfig.GetConfigId(), err.Error())
			}
		}
	}()
	if msg == "" {
		return
	}
	w.sendMsg(w.iConfig.GetConfigId(), msg)
}

func (w *worker) sendMsg(configId, msg string) {
	list, err := w.getNotifyList(w.iConfig.GetConfigId())
	if err != nil {
		log.Error(fmt.Sprintf("获取通知列表时发生错误:%s，消息未发送：%s", err.Error(), msg))
		return
	}
	sendFlag := false
	for _, n := range list {
		err = n.SendMsg(msg)
		if err != nil {
			log.Error(fmt.Sprintf("发送消息时遇到错误:%s，消息未发送：%s", err.Error(), msg))
		} else {
			sendFlag = true
		}
	}
	if !sendFlag {
		log.Warn(fmt.Sprintf("消息未发送：%s", msg))
	}
}

func (w *worker) getNotifyList(id string) ([]notify.INotify, error) {
	nl := repository.NotifyList{}
	d, err := nl.GetNotifyList(id)
	if err != nil {
		errMsg := fmt.Sprintf("获取通知配置时发生错误：%s", err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	//result := make([]notify.INotify,0)
	//dingTalk := repository.NewNotifyRepository(&repository.DingTalkRobotConfig{})
	//for _,id := range d.DingTalkRobot {
	//	nd,err := dingTalk.GetNotify(id)
	//	if err != nil {
	//		return nil,err
	//	}
	//	notify.NewNotify()
	//
	//}

	//TODO
	return nil, nil
}
