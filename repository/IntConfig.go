package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetIntTaskConfig = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]"

const SqlGetIntTaskConfigById = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]" +
	" WHERE B.[FId]=?"

type IntConfig struct {
}

type IntConfigData struct {
	FId         string
	FServer     string
	FPort       int
	FDbName     string
	FDbUser     string
	FDbPwd      string
	FSearch     string
	FCron       string
	FCheckMax   int
	FCheckMin   int
	FMsgTitle   string
	FMsgContent string
}

func (configData *IntConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *IntConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *IntConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case IntConfigData:
		if configData.FId != c.FId ||
			configData.FServer != c.FServer ||
			configData.FPort != c.FPort ||
			configData.FDbName != c.FDbName ||
			configData.FDbUser != c.FDbUser ||
			configData.FDbPwd != c.FDbPwd ||
			configData.FSearch != c.FSearch ||
			configData.FCron != c.FCron ||
			configData.FCheckMax != c.FCheckMax ||
			configData.FCheckMin != c.FCheckMin ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntConfigData"))
		return false
	}
}

func (ic *IntConfig) GetSqlGetConfigList() string {
	return SqlGetIntTaskConfig
}

func (ic *IntConfig) GetSqlGetConfig() string {
	return SqlGetIntTaskConfigById
}

func (ic *IntConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(
			&fId, &fServer, &fPort, &fDbName, &fDbUser,
			&fDbPwd, &fSearch, &fCron, &fCheckMax, &fCheckMin,
			&fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := IntConfigData{
			FId:         fId,
			FServer:     fServer,
			FPort:       fPort,
			FDbName:     fDbName,
			FDbUser:     fDbUser,
			FDbPwd:      fDbPwd,
			FSearch:     fSearch,
			FCron:       fCron,
			FCheckMax:   fCheckMax,
			FCheckMin:   fCheckMin,
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
