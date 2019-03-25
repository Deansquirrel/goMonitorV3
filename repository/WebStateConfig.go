package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetWebStateTaskConfig = "" +
	"SELECT B.[FId], B.[FUrl], B.[FCron], B.[FMsgTitle], B.[FMsgContent]" +
	" FROM MConfig A" +
	" INNER JOIN WebStateTaskConfig B ON A.FID = B.FId"

const SqlGetWebStateTaskConfigById = "" +
	"SELECT B.[FId], B.[FUrl], B.[FCron], B.[FMsgTitle], B.[FMsgContent]" +
	" FROM MConfig A" +
	" INNER JOIN WebStateTaskConfig B ON A.FID = B.FId" +
	" WHERE B.FId = ?"

type webStateConfig struct {
}

func (wsc *webStateConfig) GetSqlGetConfigList() string {
	return SqlGetWebStateTaskConfig
}

func (wsc *webStateConfig) GetSqlGetConfig() string {
	return SqlGetWebStateTaskConfigById
}

func (wsc *webStateConfig) getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fUrl, fCron, fMsgTitle, fMsgContent string
	resultList := make([]object.IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fUrl, &fCron, &fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := object.WebStateConfigData{
			FId:         fId,
			FUrl:        fUrl,
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
