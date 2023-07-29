package queue

type Queue interface {
	ReceiveMessages(num int) ([]interface{}, error)
	SendMessages(messages ...interface{}) error
}
