package taskService

import (
	"reflect"
)

type intTask struct {
}

func (it *intTask) getCacheIdList() []string {
	list := make([]string, 0)
	for _, c := range GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*repository.IntConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
