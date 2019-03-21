package notify

type INotify interface {
	SendMsg(msg string) error
}
