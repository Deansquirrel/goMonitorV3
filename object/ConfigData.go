package object

import (
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
)

type MConfigData struct {
	FId     string
	FTitle  string
	FRemark string
}

func (configData *MConfigData) GetSpec() string {
	return ""
}

func (configData *MConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *MConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.mConfigData":
		c, ok := d.(*MConfigData)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FTitle != c.FTitle ||
			configData.FRemark != c.FRemark {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：mConfigData"))
		return false
	}
}

type IntConfigData struct {
	FId         string
	FServer     string
	FPort       int
	FDbName     string
	FDbUser     string
	FDbPwd      string
	FSearch     string
	FCron       string
	FCheckMax   int
	FCheckMin   int
	FMsgTitle   string
	FMsgContent string
}

func (configData *IntConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *IntConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *IntConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.IntConfigData":
		c, ok := d.(*IntConfigData)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FServer != c.FServer ||
			configData.FPort != c.FPort ||
			configData.FDbName != c.FDbName ||
			configData.FDbUser != c.FDbUser ||
			configData.FDbPwd != c.FDbPwd ||
			configData.FSearch != c.FSearch ||
			configData.FCron != c.FCron ||
			configData.FCheckMax != c.FCheckMax ||
			configData.FCheckMin != c.FCheckMin ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntConfigData"))
		return false
	}
}

type IntDConfigData struct {
	FId        string
	FMsgSearch string
}

func (configData *IntDConfigData) GetSpec() string {
	return ""
}

func (configData *IntDConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *IntDConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.IntDConfigData":
		c, ok := d.(*IntDConfigData)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FMsgSearch != c.FMsgSearch {
			return false
		}
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntDConfigData"))
		return false
	}
}

type HealthConfigData struct {
	FId         string
	FCron       string
	FMsgTitle   string
	FMsgContent string
}

func (configData *HealthConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *HealthConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *HealthConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.healthConfigData":
		c, ok := d.(*HealthConfigData)
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
		log.Warn(fmt.Sprintf("expr：healthConfigData"))
		return false
	}
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

func (configData *CrmDzXfTestConfigData) GetSpec() string {
	return configData.FCron
}

func (configData *CrmDzXfTestConfigData) GetConfigId() string {
	return configData.FId
}

func (configData *CrmDzXfTestConfigData) IsEqual(d IConfigData) bool {
	switch reflect.TypeOf(d).String() {
	case "*repository.crmDzXfTestConfigData":
		c, ok := d.(*CrmDzXfTestConfigData)
		if !ok {
			return false
		}
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
		log.Warn(fmt.Sprintf("expr：crmDzXfTestConfigData"))
		return false
	}
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
