package worker

import "github.com/Deansquirrel/goMonitorV3/repository"

type IWorker interface {
	GetMsg() (string, repository.IHisData)
	SaveSearchResult(data interface{})

	formatMsg(msg string) string
}
