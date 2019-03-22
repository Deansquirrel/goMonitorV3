package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetIntTaskDConfig = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig]"

const SqlGetIntTaskDConfigById = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig] " +
	"WHERE [FId]=?"

type IntDConfig struct {
}

type IntDConfigData struct {
	FId        string
	FMsgSearch string
}

func (configData *IntDConfigData) GetSpec() string {
	return ""
}

func (configData *IntDConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *IntDConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case IntDConfigData:
		if configData.FId != c.FId ||
			configData.FMsgSearch != c.FMsgSearch {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntDConfigData"))
		return false
	}
}

func (idc *IntDConfig) GetSqlGetConfigList() string {
	return SqlGetIntTaskDConfig
}

func (idc *IntDConfig) GetSqlGetConfig() string {
	return SqlGetIntTaskDConfigById
}

func (idc *IntDConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fMsgSearch sql.NullString
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err := rows.Scan(&fId, &fMsgSearch)
		if err != nil {
			break
		}
		config := IntDConfigData{}
		config.FId = "Null"
		if fId.Valid {
			config.FId = fId.String
		}
		config.FMsgSearch = "Null"
		if fMsgSearch.Valid {
			config.FMsgSearch = fMsgSearch.String
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
