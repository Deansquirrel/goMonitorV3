package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetTaskMConfig = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig]"
const SqlGetTaskMConfigById = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig] WHERE [FID] = ?"

type mConfig struct {
}

func (config *mConfig) GetSqlGetConfigList() string {
	return SqlGetTaskMConfig
}

func (config *mConfig) GetSqlGetConfig() string {
	return SqlGetTaskMConfigById
}

func (config *mConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fTitle, fRemark string
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fTitle, &fRemark)
		if err != nil {
			return nil, err
		}
		config := object.MConfigData{
			FId:     fId,
			FTitle:  fTitle,
			FRemark: fRemark,
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
