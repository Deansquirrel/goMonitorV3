package configRepository

import (
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetIntTaskConfig = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]"

const SqlGetIntTaskConfigById = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]" +
	" WHERE B.[FId]=?"

type IntConfig struct {
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

func (icd *IntConfigData) IsEqual(d interface{}) bool {
	switch d.(type) {
	case IntConfigData:
		return true
	default:
		log.Warn(fmt.Sprintf("expr：IntConfigData"))
		return false
	}
}

func (ic *IntConfig) GetSqlGetConfigList() string {
	return SqlGetIntTaskConfig
}

func (ic *IntConfig) GetSqlGetConfig() string {
	return SqlGetIntTaskConfigById
}

func (ic *IntConfig) GetScanWrapperList() []interface{} {
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	list := make([]interface{}, 0)
	list = append(list, fId)
	list = append(list, fServer)
	list = append(list, fPort)
	list = append(list, fDbName)
	list = append(list, fDbUser)
	list = append(list, fDbPwd)
	list = append(list, fSearch)
	list = append(list, fCron)
	list = append(list, fCheckMax)
	list = append(list, fCheckMin)
	list = append(list, fMsgTitle)
	list = append(list, fMsgContent)
	return list
}

func (ic *IntConfig) GetConfigDataByScanWrapperList(arg ...interface{}) IConfigData {
	fmt.Println(arg)
	//TODO
	config := IntConfigData{
		//FId:string(arg[0]),
		//FServer:arg[1].(string),
		//FPort:       arg[2].(int),
		//FDbName:     arg[3].(string),
		//FDbUser:     arg[4].(string),
		//FDbPwd:      arg[5].(string),
		//FSearch:     arg[6].(string),
		//FCron:       arg[7].(string),
		//FCheckMin:   arg[9].(int),
		//FMsgTitle:   arg[10].(string),
		//FMsgContent: arg[11].(string),
	}
	return &config
}

func (ic *IntConfig) CheckConfigData(configData IConfigData) bool {
	return false
}
