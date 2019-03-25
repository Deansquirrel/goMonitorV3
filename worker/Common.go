package worker

import (
	"strings"
)

type common struct {
}

//获取待发送消息
func (c *common) getMsg(title, content string) string {
	msg := ""
	titleList := strings.Split(title, "###")
	if len(titleList) > 0 {
		for _, t := range titleList {
			if strings.Trim(t, " ") != "" {
				if msg != "" {
					msg = msg + "\n"
				}
				msg = msg + strings.Trim(t, " ")
			}
		}
	}
	contentList := strings.Split(content, "###")
	if len(contentList) > 0 {
		for _, t := range contentList {
			if strings.Trim(t, " ") != "" {
				if msg != "" {
					msg = msg + "\n"
				}
				msg = msg + strings.Trim(t, " ")
			}
		}
	}
	return msg
}
