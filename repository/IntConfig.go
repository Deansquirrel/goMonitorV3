package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
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

type intConfig struct {
}

func (ic *intConfig) GetSqlGetConfigList() string {
	return SqlGetIntTaskConfig
}

func (ic *intConfig) GetSqlGetConfig() string {
	return SqlGetIntTaskConfigById
}

func (ic *intConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(
			&fId, &fServer, &fPort, &fDbName, &fDbUser,
			&fDbPwd, &fSearch, &fCron, &fCheckMax, &fCheckMin,
			&fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := object.IntConfigData{
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
