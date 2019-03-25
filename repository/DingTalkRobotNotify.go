package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
)

import log "github.com/Deansquirrel/goToolLog"

//const SqlGetDingTalkRobot = "" +
//	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
//	" FROM [NConfig] A" +
//	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]"

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

type dingTalkRobotNotify struct {
}

func (config *dingTalkRobotNotify) GetSqlGetConfig() string {
	return SqlGetDingTalkRobotById
}

func (config *dingTalkRobotNotify) getConfigListByRows(rows *sql.Rows) ([]object.INotifyData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fWebHookKey, fAtMobiles string
	var fIsAtAll int
	resultList := make([]object.INotifyData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fWebHookKey, &fAtMobiles, &fIsAtAll)
		if err != nil {
			break
		}
		config := object.DingTalkRobotNotifyData{
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
