package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetCrmDzXfTestTaskConfig = "" +
	"SELECT B.[FId],B.[FCron],B.[FMsgTitle],B.[FMsgContent],B.[FAddress],B.[FPassport],B.[FPassportType]" +
	" FROM MConfig A" +
	" INNER JOIN CrmDzXfTestTaskConfig B ON A.[FId] = B.[FId]"

const SqlGetCrmDzXfTestTaskConfigById = "" +
	"SELECT B.[FId],B.[FCron],B.[FMsgTitle],B.[FMsgContent],B.[FAddress],B.[FPassport],B.[FPassportType]" +
	" FROM MConfig A" +
	" INNER JOIN CrmDzXfTestTaskConfig B ON A.[FId] = B.[FId]" +
	" WHERE B.FId = ?"

type crmDzXfTestConfig struct {
}

func (config *crmDzXfTestConfig) GetSqlGetConfigList() string {
	return SqlGetCrmDzXfTestTaskConfig
}

func (config *crmDzXfTestConfig) GetSqlGetConfig() string {
	return SqlGetCrmDzXfTestTaskConfigById
}

func (config *crmDzXfTestConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fCron, fMsgTitle, fMsgContent, fAddress, fPassport string
	var fPassportType int
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fCron, &fMsgTitle, &fMsgContent, &fAddress, &fPassport, &fPassportType)
		if err != nil {
			break
		}
		config := object.CrmDzXfTestConfigData{
			FId:           fId,
			FCron:         fCron,
			FMsgTitle:     fMsgTitle,
			FMsgContent:   fMsgContent,
			FAddress:      fAddress,
			FPassport:     fPassport,
			FPassportType: fPassportType,
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
