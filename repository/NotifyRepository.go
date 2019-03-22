package repository

import (
	"errors"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

type notifyRepository struct {
	Config INotify
}

func NewNotifyRepository(config INotify) *notifyRepository {
	return &notifyRepository{
		Config: config,
	}
}

func (nr *notifyRepository) GetNotify(id string) (INotifyData, error) {
	rows, err := comm.getRowsBySQL(nr.Config.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	list, err := nr.Config.getConfigListByRows(rows)
	if err != nil {
		return nil, err
	}
	if len(list) < 1 {
		return nil, nil
	}
	if len(list) > 1 {
		return nil, errors.New(fmt.Sprintf("Config列表数量异常，exp：1，act:%d", len(list)))
	}
	return list[0], nil
}
