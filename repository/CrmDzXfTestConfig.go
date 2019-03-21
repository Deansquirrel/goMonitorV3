package repository

import (
	"database/sql"
	"fmt"
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

type CrmDzXfTestConfig struct {
}

type CrmDzXfTestConfigData struct {
	FId           string
	FCron         string
	FMsgTitle     string
	FMsgContent   string
	FAddress      string
	FPassport     string
	FPassportType int
}

func (configData *CrmDzXfTestConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case CrmDzXfTestConfigData:
		if configData.FId != c.FId ||
			configData.FCron != c.FCron ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent ||
			configData.FAddress != c.FAddress ||
			configData.FPassport != c.FPassport ||
			configData.FPassportType != c.FPassportType {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：CrmDzXfTestConfigData"))
		return false
	}
}

func (config *CrmDzXfTestConfig) GetSqlGetConfigList() string {
	return SqlGetCrmDzXfTestTaskConfig
}

func (config *CrmDzXfTestConfig) GetSqlGetConfig() string {
	return SqlGetCrmDzXfTestTaskConfigById
}

func (config *CrmDzXfTestConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fCron, fMsgTitle, fMsgContent, fAddress, fPassport string
	var fPassportType int
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fCron, &fMsgTitle, &fMsgContent, &fAddress, &fPassport, &fPassportType)
		if err != nil {
			break
		}
		config := CrmDzXfTestConfigData{
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
