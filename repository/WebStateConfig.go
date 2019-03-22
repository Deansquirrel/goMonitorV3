package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
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

func (configData *WebStateConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *WebStateConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *WebStateConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.WebStateConfigData":
		c, ok := d.(*WebStateConfigData)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FCron != c.FCron ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent {
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
