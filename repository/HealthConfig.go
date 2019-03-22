package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
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

type HealthConfig struct {
}

type HealthConfigData struct {
	FId         string
	FCron       string
	FMsgTitle   string
	FMsgContent string
}

func (configData *HealthConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *HealthConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *HealthConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.HealthConfigData":
		c, ok := d.(*HealthConfigData)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FCron != c.FCron ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：HealthConfigData"))
		return false
	}
}

func (hc *HealthConfig) GetSqlGetConfigList() string {
	return SqlGetHealthConfig
}

func (hc *HealthConfig) GetSqlGetConfig() string {
	return SqlGetHealthConfigById
}

func (hc *HealthConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fCron, fMsgTitle, fMsgContent string
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fCron, &fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := HealthConfigData{
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
