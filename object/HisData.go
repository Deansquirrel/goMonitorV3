package object

import "time"

type CrmDzXfTestHisData struct {
	FId       string
	FConfigId string
	FUseTime  int
	FHttpCode int
	FContent  string
	FOprTime  time.Time
}

type IntHisData struct {
	FId       string
	FConfigId string
	FNum      int
	FContent  string
	FOprTime  time.Time
}

type WebStateHisData struct {
	FId       string
	FConfigId string
	FUseTime  int
	FHttpCode int
	FContent  string
	FOprTime  time.Time
}
