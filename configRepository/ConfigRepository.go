package configRepository

import (
	"database/sql"
	log "github.com/Deansquirrel/goToolLog"
)

type configRepository struct {
	Config     IConfig
	ConfigData IConfigData
}

func NewConfigRepository(config IConfig) *configRepository {
	return &configRepository{
		Config: config,
	}
}

//func NewConfigRepositoryT(config IConfig,configData IConfigData)(*configRepository,error){
//	if config.CheckConfigData(configData) {
//		return &configRepository{
//			Config:config,
//			ConfigData:configData,
//		},nil
//	} else {
//		return nil,errors.New("不匹配的配置类型")
//	}
//}

func (cr *configRepository) GetConfigList() ([]interface{}, error) {
	rows, err := comm.getRowsBySQL(cr.Config.GetSqlGetConfigList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cr.getConfigListByRows(rows)
}

func (cr *configRepository) GetConfig(id string) (interface{}, error) {
	rows, err := comm.getRowsBySQL(cr.Config.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cr.getConfigListByRows(rows)
}

func (cr *configRepository) getConfigListByRows(rows *sql.Rows) ([]interface{}, error) {
	defer func() {
		_ = rows.Close()
	}()
	resultList := make([]interface{}, 0)
	var err error
	swList := cr.Config.GetScanWrapperList()
	swPList := make([]interface{}, 0)
	for i := 0; i < len(swList); i++ {
		swPList = append(swPList, &swList[i])
	}
	for rows.Next() {
		err = rows.Scan(swPList...)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		configData := cr.Config.GetConfigDataByScanWrapperList(swList...)
		resultList = append(resultList, configData)
	}
	if rows.Err() != nil {
		log.Error(rows.Err().Error())
		return nil, rows.Err()
	}
	return resultList, nil
}
