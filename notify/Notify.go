package notify

import (
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/kataras/iris/core/errors"
	"reflect"
)

type notify struct {
	iNotify repository.INotifyData
}

func NewNotify(iNotify repository.INotifyData) *notify {
	return &notify{
		iNotify: iNotify,
	}
}

//func (n *notify) GetNotify(id string)(INotify,error){
//	rep := repository.NewNotifyRepository(n.iNotify)
//	iNotify,err := rep.GetNotify(id)
//	if err != nil {
//		return nil,err
//	}
//	return n.getNotify(iNotify)
//}

func (n *notify) GetNotify(iNotify repository.INotifyData) (INotify, error) {
	switch reflect.TypeOf(iNotify).String() {
	case "*repository.DingTalkRobotConfigData":
		dingTalkRobotConfigData, ok := iNotify.(*repository.DingTalkRobotConfigData)
		if ok {
			return newDingTalkRobot(dingTalkRobotConfigData), nil
		} else {
			return nil, errors.New("强制类型转换失败[DingTalkRobotConfigData]")
		}
	default:
		return nil, errors.New("未预知的通知类型")
	}
}
