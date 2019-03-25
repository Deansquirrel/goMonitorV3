package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetHealthConfig = "" +
	"SELECT B.FId,B.FCron,B.FMsgTitle,B.FMsgContent" +
	" From MConfig A" +
	" INNER JOIN HealthTaskConfig B on A.FId = B.FId"

const SqlGetHealthConfigById = "" +
	"SELECT B.FId,B.FCron,B.FMsgTitle,B.FMsgContent" +
	" From MConfig A" +
	" INNER JOIN HealthTaskConfig B on A.FId = B.FId" +
	" WHERE FId = ?"

type healthConfig struct {
}

func (hc *healthConfig) GetSqlGetConfigList() string {
	return SqlGetHealthConfig
}

func (hc *healthConfig) GetSqlGetConfig() string {
	return SqlGetHealthConfigById
}

func (hc *healthConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fCron, fMsgTitle, fMsgContent string
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fCron, &fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := object.HealthConfigData{
			FId:         fId,
			FCron:       fCron,
			FMsgTitle:   fMsgTitle,
			FMsgContent: fMsgContent,
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
