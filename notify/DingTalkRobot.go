package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/global"
	"github.com/Deansquirrel/goMonitorV3/object"
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"io/ioutil"
	"net/http"
	"strings"
)

type dingTalkRobot struct {
	configData *repository.DingTalkRobotConfigData
}

type dingTalkTextMsg struct {
	WebHookKey string   `json:"webhookkey"`
	Content    string   `json:"content"`
	AtMobiles  []string `json:"atmobiles"`
	IsAtAll    bool     `json:"isatall"`
}

func newDingTalkRobot(configData *repository.DingTalkRobotConfigData) *dingTalkRobot {
	return &dingTalkRobot{
		configData: configData,
	}
}

func (dt *dingTalkRobot) SendMsg(msg string) error {
	var err error
	if dt.configData.FIsAtAll == 1 {
		err = dt.sendTextMsgWithAtAll(dt.configData.FWebHookKey, msg)
	} else if strings.Trim(dt.configData.FAtMobiles, " ") != "" {
		list := strings.Split(strings.Trim(dt.configData.FAtMobiles, " "), ",")
		list = goToolCommon.ClearBlock(list)
		if len(list) > 0 {
			err = dt.sendTextMsgWithAtList(dt.configData.FWebHookKey, msg, list)
		} else {
			err = dt.sendTextMsg(dt.configData.FWebHookKey, msg)
		}
	}
	return err
}

func (dt *dingTalkRobot) sendTextMsg(webHookKey string, msg string) error {
	om := dingTalkTextMsg{
		WebHookKey: webHookKey,
		Content:    msg,
	}
	return dt.sendMsg(om)
}

func (dt *dingTalkRobot) sendTextMsgWithAtList(webHookKey string, msg string, atMobiles []string) error {
	om := dingTalkTextMsg{
		WebHookKey: webHookKey,
		Content:    msg,
		AtMobiles:  atMobiles,
	}
	return dt.sendMsg(om)
}

func (dt *dingTalkRobot) sendTextMsgWithAtAll(webHookKey string, msg string) error {
	om := dingTalkTextMsg{
		WebHookKey: webHookKey,
		Content:    msg,
		IsAtAll:    true,
	}
	return dt.sendMsg(om)
}

//获取Text消息发送地址
func (dt *dingTalkRobot) getTextMsgUrl() string {
	return fmt.Sprintf("%s/text", global.SysConfig.DingTalkConfig.Address)
}

//发送消息
func (dt *dingTalkRobot) sendMsg(v interface{}) error {
	msg, err := goToolCommon.GetJsonStr(v)
	if err != nil {
		return err
	}
	rData, err := dt.sendData([]byte(msg), dt.getTextMsgUrl())
	if err != nil {
		return err
	}
	return dt.tranRData(rData)
}

//POST发送数据
func (dt *dingTalkRobot) sendData(data []byte, url string) ([]byte, error) {
	log.Debug(string(data))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("构造http请求数据时发生错误：" + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("发送http请求时错误：" + err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	rData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取http返回数据时发生错误：" + err.Error())
	}
	return rData, nil
}

//检查返回数据
func (dt *dingTalkRobot) tranRData(data []byte) error {
	var r object.SimpleResponse
	err := json.Unmarshal(data, &r)
	if err != nil {
		return errors.New(fmt.Sprintf("返回数据格式化异常，err：[%s]，返回数据：[%s]", err.Error(), string(data)))
	}
	if r.ErrCode == 0 && strings.ToLower(r.ErrMsg) == "ok" {
		return nil
	} else {
		if strings.Trim(r.ErrMsg, " ") != "" {
			return errors.New(r.ErrMsg)
		} else {
			return errors.New(fmt.Sprintf("未知错误，errcode[%d]", r.ErrCode))
		}
	}
}
