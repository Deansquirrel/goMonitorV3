package repository

import (
	"database/sql"
	"fmt"
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

type WebStateConfig struct {
}

type WebStateConfigData struct {
	FId         string
	FUrl        string
	FCron       string
	FMsgTitle   string
	FMsgContent string
}

func (config *WebStateConfigData) IsEqual(d interface{}) bool {
	switch c := d.(type) {
	case WebStateConfigData:
		if config.FId != c.FId ||
			config.FCron != c.FCron ||
			config.FMsgTitle != c.FMsgTitle ||
			config.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：WebStateConfigData"))
		return false
	}
}

func (wsc *WebStateConfig) GetSqlGetConfigList() string {
	return SqlGetWebStateTaskConfig
}

func (wsc *WebStateConfig) GetSqlGetConfig() string {
	return SqlGetWebStateTaskConfigById
}

func (wsc *WebStateConfig) getConfigListByRows(rows *sql.Rows) ([]IConfigData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fUrl, fCron, fMsgTitle, fMsgContent string
	resultList := make([]IConfigData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fUrl, &fCron, &fMsgTitle, &fMsgContent)
		if err != nil {
			break
		}
		config := WebStateConfigData{
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
