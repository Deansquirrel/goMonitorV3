package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetTaskMConfig = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig]"
const SqlGetTaskMConfigById = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig] WHERE [FID] = ?"

type MConfig struct {
}

type MConfigData struct {
	FId     string
	FTitle  string
	FRemark string
}

func (configData *MConfigData) GetSpec() string {
	return ""
}

func (configData *MConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *MConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case MConfigData:
		if configData.FId != c.FId ||
			configData.FTitle != c.FTitle ||
			configData.FRemark != c.FRemark {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：MConfigData"))
		return false
	}
}

func (config *MConfig) GetSqlGetConfigList() string {
	return SqlGetTaskMConfig
}

func (config *MConfig) GetSqlGetConfig() string {
	return SqlGetTaskMConfigById
}

func (config *MConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fTitle, fRemark string
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fTitle, &fRemark)
		if err != nil {
			return nil, err
		}
		config := MConfigData{
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
