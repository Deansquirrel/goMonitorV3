package taskService

import (
	"reflect"
)

type IntTask struct {
}

func (it *IntTask) getCacheIdList() []string {
	list := make([]string, 0)
	for _, c := range GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*repository.IntConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
