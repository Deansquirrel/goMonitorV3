package object

type IConfigData interface {
	GetConfigId() string
	GetSpec() string
	IsEqual(c IConfigData) bool
}
