package worker

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/notify"
	"github.com/Deansquirrel/goMonitorV3/repository"
	log "github.com/Deansquirrel/goToolLog"
)

type worker struct {
	configId string
	config   interface{}
	iWorker  IWorker
}

func (w *worker) NewWorker(config interface{}) (*worker, error) {
	id, workerRunner, err := w.getWorker(config)
	if err != nil {
		return nil, err
	}
	return &worker{
		configId: id,
		config:   config,
		iWorker:  workerRunner,
	}, nil
}

func (w *worker) Run() {
	msg, hisData := w.iWorker.GetMsg()
	defer func() {
		if hisData != nil {
			w.iWorker.SaveSearchResult(hisData)
		}
	}()
	if msg == "" {
		return
	}
	list, err := w.getNotifyList(w.configId)
	if err != nil {
		log.Error(fmt.Sprintf("获取通知列表时发生错误:%s，消息未发送：%s", err.Error(), msg))
		return
	}
	for _, n := range list {
		err = n.SendMsg(msg)
		if err != nil {
			log.Error(fmt.Sprintf("发送消息时遇到错误:%s，消息未发送：%s", err.Error(), msg))
		}
	}
}

func (w *worker) getWorker(config interface{}) (string, IWorker, error) {
	switch c := config.(type) {
	case repository.CrmDzXfTestConfigData:
		return "", nil, nil
	case repository.HealthConfigData:
		return "", nil, nil
	case repository.IntConfigData:
		return c.FId, NewIntWorker(&c), nil
	case repository.WebStateConfigData:
		return "", nil, nil
	default:
		return "", nil, errors.New("未预知的配置类型")
	}
}

func (w *worker) getNotifyList(id string) ([]notify.INotify, error) {
	return nil, nil
}
