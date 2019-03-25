package notify

import (
	"github.com/Deansquirrel/goMonitorV3/object"
	"github.com/kataras/iris/core/errors"
	"reflect"
)

type notify struct {
}

func NewNotify() *notify {
	return &notify{}
}

func (n *notify) GetNotify(iNotify object.INotifyData) (INotify, error) {
	switch reflect.TypeOf(iNotify).String() {
	case "*object.DingTalkRobotConfigData":
		dingTalkRobotConfigData, ok := iNotify.(*object.DingTalkRobotNotifyData)
		if ok {
			return newDingTalkRobot(dingTalkRobotConfigData), nil
		} else {
			return nil, errors.New("强制类型转换失败[DingTalkRobotConfigData]")
		}
	default:
		return nil, errors.New("未预知的通知类型")
	}
}
