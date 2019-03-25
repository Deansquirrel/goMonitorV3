package worker

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/notify"
	"github.com/Deansquirrel/goMonitorV3/object"
	"github.com/Deansquirrel/goMonitorV3/repository"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
)

type worker struct {
	iConfig object.IConfigData
	iWorker IWorker
}

func NewWorker(iConfig object.IConfigData) (*worker, error) {
	workerRunner, err := getWorker(iConfig)
	if err != nil {
		return nil, err
	}
	return &worker{
		iConfig: iConfig,
		iWorker: workerRunner,
	}, nil
}

func getWorker(iConfig object.IConfigData) (IWorker, error) {
	switch reflect.TypeOf(iConfig).String() {
	case "*object.IntConfigData":
		config, ok := iConfig.(*object.IntConfigData)
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

	notifyDataList := make([]object.INotifyData, 0)

	errMsg := ""
	errMsgFormat := "获取[%s]NotifyData时发生错误：%s；"

	//获取DingTalkRobot类型配置数据
	dingTalkRep := repository.NewDingTalkRobotNotifyRepository()
	for _, id := range d.DingTalkRobot {
		n, err := dingTalkRep.GetNotify(id)
		if err != nil {
			m := fmt.Sprintf(errMsgFormat, "DingTalkRobot", err.Error())
			log.Error(m)
			errMsg = errMsg + m
		} else {
			notifyDataList = append(notifyDataList, n)
		}
	}

	errMsgFormat = "获取Notify时发生错误：%s,NotifyID:%s"
	result := make([]notify.INotify, 0)
	ns := notify.NewNotify()
	for _, nd := range notifyDataList {
		n, err := ns.GetNotify(nd)
		if err != nil {
			m := fmt.Sprintf(errMsgFormat, err.Error(), nd.GetNotifyId())
			errMsg = errMsg + m
		} else {
			result = append(result, n)
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	} else {
		err = nil
	}

	return result, err
}
