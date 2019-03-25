package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetIntTaskDConfig = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig]"

const SqlGetIntTaskDConfigById = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig] " +
	"WHERE [FId]=?"

type intDConfig struct {
}

func (idc *intDConfig) GetSqlGetConfigList() string {
	return SqlGetIntTaskDConfig
}

func (idc *intDConfig) GetSqlGetConfig() string {
	return SqlGetIntTaskDConfigById
}

func (idc *intDConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fMsgSearch sql.NullString
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err := rows.Scan(&fId, &fMsgSearch)
		if err != nil {
			break
		}
		config := object.IntDConfigData{}
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
