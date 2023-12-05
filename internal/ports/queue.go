package ports

type Queue interface {
	Connect() error
	Publish(queueName string, data []byte) error
	Consume(queueName string, callback func(args []byte) error) error
}
