package repository

import (
	"database/sql"
	"fmt"
)

import log "github.com/Deansquirrel/goToolLog"

const SqlGetDingTalkRobot = "" +
	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
	" FROM [NConfig] A" +
	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]"

const SqlGetDingTalkRobotById = "" +
	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
	" FROM [NConfig] A" +
	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]" +
	" WHERE A.[FId]=?"

//const SqlGetDingTalkRobotByIdList = "" +
//	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
//	" FROM [NConfig] A" +
//	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]" +
//	" WHERE A.[FId] in (%s)"

type DingTalkRobotConfig struct {
}

type DingTalkRobotConfigData struct {
	FId         string
	FWebHookKey string
	FAtMobiles  string
	FIsAtAll    int
}

func (configData *DingTalkRobotConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case DingTalkRobotConfigData:
		if configData.FId != c.FId ||
			configData.FWebHookKey != c.FWebHookKey ||
			configData.FAtMobiles != c.FAtMobiles ||
			configData.FIsAtAll != c.FIsAtAll {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntDConfigData"))
		return false
	}
}

func (config *DingTalkRobotConfig) GetSqlGetConfigList() string {
	return SqlGetDingTalkRobot
}

func (config *DingTalkRobotConfig) GetSqlGetConfig() string {
	return SqlGetDingTalkRobotById
}

func (config *DingTalkRobotConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fWebHookKey, fAtMobiles string
	var fIsAtAll int
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fWebHookKey, &fAtMobiles, &fIsAtAll)
		if err != nil {
			break
		}
		config := DingTalkRobotConfigData{
			FId:         fId,
			FWebHookKey: fWebHookKey,
			FAtMobiles:  fAtMobiles,
			FIsAtAll:    fIsAtAll,
		}
		resultList = append(resultList, &config)
	}
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if rows.Err() != nil {
		log.Error(rows.Err().Error())
		return nil, rows.Err()
	}
	return resultList, nil
}
