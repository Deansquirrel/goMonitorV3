package configRepository

type IConfig interface {
	GetSqlGetConfigList() string
	GetSqlGetConfig() string

	GetScanWrapperList() []interface{}
	GetConfigDataByScanWrapperList(arg ...interface{}) IConfigData

	CheckConfigData(configData IConfigData) bool
}

type IConfigData interface {
	IsEqual(c interface{}) bool
}
