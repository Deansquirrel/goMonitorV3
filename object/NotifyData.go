package object

type DingTalkRobotNotifyData struct {
	FId         string
	FWebHookKey string
	FAtMobiles  string
	FIsAtAll    int
}

func (d *DingTalkRobotNotifyData) GetNotifyId() string {
	return d.FId
}
