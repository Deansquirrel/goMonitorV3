package repository

import (
	"database/sql"
	"errors"
	"github.com/Deansquirrel/goMonitorV3/object"
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

type intHis struct {
}

func (ih *intHis) GetSqlHisList() string {
	return SqlGetIntTaskHis
}

func (ih *intHis) GetSqlHisById() string {
	return SqlGetIntTaskHisById
}

func (ih *intHis) GetSqlHisByConfigId() string {
	return SqlGetIntTaskHisByConfigId
}

func (ih *intHis) GetSqlHisByTime() string {
	return SqlGetIntTaskHisByTime
}

func (ih *intHis) GetSqlSetHis() string {
	return SqlSetIntTaskHis
}
func (ih *intHis) GetSqlClearHis() string {
	return SqlDelIntTaskHisByOprTime
}

func (ih *intHis) getHisListByRows(rows *sql.Rows) ([]object.IHisData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fConfigId, fContent string
	var fNum int
	var fOprTime time.Time
	resultList := make([]object.IHisData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fConfigId, &fNum, &fContent, &fOprTime)
		if err != nil {
			break
		}
		config := object.IntHisData{
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

func (ih *intHis) getHisSetArgs(data object.IHisData) ([]interface{}, error) {
	switch reflect.TypeOf(data).String() {
	case "*object.intHisData":
		iHisData, ok := data.(*object.IntHisData)
		if ok {
			result := make([]interface{}, 0)
			result = append(result, iHisData.FId)
			result = append(result, iHisData.FConfigId)
			result = append(result, iHisData.FNum)
			result = append(result, iHisData.FContent)
			return result, nil
		} else {
			return nil, errors.New("强制类型转换失败[intHisData]")
		}
	default:
		return nil, errors.New("intHis getHisSetArgs 参数类型错误")
	}
}
