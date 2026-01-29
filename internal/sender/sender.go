package sender

type MessageSender interface {
	Send(to string, message string) error
}
