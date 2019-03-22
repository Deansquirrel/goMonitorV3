package repository

import (
	"database/sql"
	"errors"
	"reflect"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const SqlGetIntTaskHis = "" +
	"SELECT [FId],[FConfigId],[FNum],[FContent],[FOprTime]" +
	" FROM [IntTaskHis]" +
	" Order By [FOprTime] Asc"

const SqlGetIntTaskHisById = "" +
	"SELECT [FId],[FConfigId],[FNum],[FContent],[FOprTime]" +
	" FROM [IntTaskHis]" +
	" WHERE [FId]=?" +
	" Order By [FOprTime] Asc"

const SqlGetIntTaskHisByConfigId = "" +
	"SELECT [FId],[FConfigId],[FNum],[FContent],[FOprTime]" +
	" FROM [IntTaskHis]" +
	" WHERE [FConfigId] = ?" +
	" Order By [FOprTime] Asc"

const SqlGetIntTaskHisByTime = "" +
	"SELECT [FId],[FConfigId],[FNum],[FContent],[FOprTime]" +
	" FROM [IntTaskHis]" +
	" WHERE [FOprTime] >= ? AND [FOprTime] <= ?" +
	" Order By [FOprTime] Asc"

const SqlSetIntTaskHis = "" +
	"INSERT INTO [IntTaskHis]([FId],[FConfigId],[FNum],[FContent])" +
	" SELECT ?,?,?,?"

const SqlDelIntTaskHisByOprTime = "" +
	"DELETE FROM [IntTaskHis]" +
	" WHERE [FOprTime] < ?"

type IntHis struct {
}

type IntHisData struct {
	FId       string
	FConfigId string
	FNum      int
	FContent  string
	FOprTime  time.Time
}

func (ih *IntHis) GetSqlHisList() string {
	return SqlGetIntTaskHis
}

func (ih *IntHis) GetSqlHisById() string {
	return SqlGetIntTaskHisById
}

func (ih *IntHis) GetSqlHisByConfigId() string {
	return SqlGetIntTaskHisByConfigId
}

func (ih *IntHis) GetSqlHisByTime() string {
	return SqlGetIntTaskHisByTime
}

func (ih *IntHis) GetSqlSetHis() string {
	return SqlSetIntTaskHis
}
func (ih *IntHis) GetSqlClearHis() string {
	return SqlDelIntTaskHisByOprTime
}

func (ih *IntHis) getHisListByRows(rows *sql.Rows) ([]IHisData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fConfigId, fContent string
	var fNum int
	var fOprTime time.Time
	resultList := make([]IHisData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fConfigId, &fNum, &fContent, &fOprTime)
		if err != nil {
			break
		}
		config := IntHisData{
			FId:       fId,
			FConfigId: fConfigId,
			FNum:      fNum,
			FContent:  fContent,
			FOprTime:  fOprTime,
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

func (ih *IntHis) getHisSetArgs(data IHisData) ([]interface{}, error) {
	switch reflect.TypeOf(data).String() {
	case "*repository.IntHisData":
		iHisData, ok := data.(*IntHisData)
		if ok {
			result := make([]interface{}, 0)
			result = append(result, iHisData.FId)
			result = append(result, iHisData.FConfigId)
			result = append(result, iHisData.FNum)
			result = append(result, iHisData.FContent)
			return result, nil
		} else {
			return nil, errors.New("强制类型转换失败[IntHisData]")
		}
	default:
		return nil, errors.New("IntHis getHisSetArgs 参数类型错误")
	}
}
