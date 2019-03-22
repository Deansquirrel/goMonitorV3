package worker

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV3/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"strconv"
	"strings"
	"time"
)

type intWorker struct {
	intConfigData *repository.IntConfigData
}

func newIntWorker(intConfigData *repository.IntConfigData) *intWorker {
	return &intWorker{
		intConfigData: intConfigData,
	}
}

func (iw *intWorker) GetMsg() (string, repository.IHisData) {
	comm := common{}
	if iw.intConfigData == nil {
		msg := comm.getMsg(iw.intConfigData.FMsgTitle, "配置内容为空")
		msg = iw.formatMsg(msg)
		return msg, iw.getHisData(0, msg)
	}
	num, err := iw.getCheckNum()
	if err != nil {
		msg := comm.getMsg(iw.intConfigData.FMsgTitle, "获取数量时遇到错误")
		msg = iw.formatMsg(msg)
		return msg, iw.getHisData(0, msg)
	}
	var msg string
	if num >= iw.intConfigData.FCheckMax || num <= iw.intConfigData.FCheckMin {
		msg = comm.getMsg(iw.intConfigData.FMsgTitle,
			strings.Replace(iw.intConfigData.FMsgContent, "title", strconv.Itoa(num), -1))
		dMsg := iw.getDMsg()
		if dMsg != "" {
			msg = msg + "\n" + dMsg
		}
		msg = iw.formatMsg(msg)
	}
	return msg, iw.getHisData(num, msg)
}

func (iw *intWorker) getHisData(num int, msg string) repository.IHisData {
	return &repository.IntHisData{
		FId:       strings.ToUpper(goToolCommon.Guid()),
		FConfigId: iw.intConfigData.FId,
		FNum:      num,
		FContent:  msg,
		FOprTime:  time.Now(),
	}
}

func (iw *intWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}

func (iw *intWorker) getDMsg() string {
	rep := repository.NewConfigRepository(&repository.IntDConfig{})
	dConfig, err := rep.GetConfig(iw.intConfigData.FId)
	if err != nil {
		log.Error(fmt.Sprintf("获取明细配置时遇到错误：%s，查询ID为：%s", err.Error(), iw.intConfigData.FId))
		return ""
	}
	//无明细配置
	if dConfig == nil {
		return ""
	}
	var v interface{}
	v = dConfig
	switch c := v.(type) {
	case repository.IntDConfigData:
		return iw.getSingleDMsg(c.FMsgSearch)
	default:
		log.Error(fmt.Sprintf("获取的明细配置类型异常，expr：IntDConfigData"))
		return ""
	}
}

func (iw *intWorker) getSingleDMsg(search string) string {
	if search == "" {
		return ""
	}
	rows, err := iw.getRowsBySQL(search)
	if err != nil {
		return fmt.Sprintf("查询明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	defer func() {
		_ = rows.Close()
	}()
	titleList, err := rows.Columns()
	if err != nil {
		return fmt.Sprintf("获取明细内容表头时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	counter := len(titleList)
	values := make([]interface{}, counter)
	valuePointers := make([]interface{}, counter)
	for i := 0; i < counter; i++ {
		valuePointers[i] = &values[i]
	}

	var result string
	for rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
		}
		if result != "" {
			result = result + "\n" + "--------------------"
		}
		for i := 0; i < counter; i++ {
			if result != "" {
				result = result + "\n"
			}
			var v string
			if values[i] == nil {
				v = "Null"
			} else {
				v = goToolCommon.ConvertToString(values[i])
			}
			result = result + titleList[i] + " - " + v
		}
	}
	if rows.Err() != nil {
		return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	return result
}

func (iw *intWorker) SaveSearchResult(data interface{}) {
	switch d := data.(type) {
	case repository.IntHisData:
		rep := repository.NewHisRepository(&repository.IntHis{})
		err := rep.SetHis(d)
		if err != nil {
			s, _ := goToolCommon.GetJsonStr(data)
			log.Error(fmt.Sprintf("保存查询结果时遇到错误：%s，待保存内容：%s", err.Error(), s))
		}
		return
	default:
		s, _ := goToolCommon.GetJsonStr(data)
		log.Error(fmt.Sprintf("历史数据类型异常，exp：IntHisData，待保存内容：%s", s))
		return
	}
}

//获取待检测值
func (iw *intWorker) getCheckNum() (int, error) {
	rows, err := iw.getRowsBySQL(iw.intConfigData.FSearch)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
	}()
	list := make([]int, 0)
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			log.Error(err.Error())
			break
		} else {
			list = append(list, num)
		}
	}
	if err != nil {
		return 0, err
	}
	if len(list) != 1 {
		errMsg := fmt.Sprintf("SQL返回数量异常，exp:1,act:%d", len(list))
		log.Error(errMsg)
		return 0, errors.New(errMsg)
	}
	return list[0], nil
}

//查询数据
func (iw *intWorker) getRowsBySQL(sql string) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(iw.getDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rows, err := conn.Query(sql)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return rows, nil
}

//获取DB配置
func (iw *intWorker) getDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: iw.intConfigData.FServer,
		Port:   iw.intConfigData.FPort,
		DbName: iw.intConfigData.FDbName,
		User:   iw.intConfigData.FDbUser,
		Pwd:    iw.intConfigData.FDbPwd,
	}
}
