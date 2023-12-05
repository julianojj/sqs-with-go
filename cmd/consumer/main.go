package main

import (
	"app/internal/adapters"
	"app/internal/application/usecases"
	"app/internal/ports"
	"encoding/json"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	sqs := &adapters.Sqs{}
	err := sqs.Connect()
	if err != nil {
		panic("error to connect to sqs")
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With(
		"service", os.Getenv("SERVICE_NAME_CONSUMER"),
		"environment", os.Getenv("SERVICE_ENV"),
	)
	notifyCustomer := usecases.NewNotifyCustomer(logger)
	Worker(sqs, notifyCustomer, logger)
	forever := make(chan bool)
	<-forever
}

func Worker(
	queue ports.Queue,
	notifyUser *usecases.NotifyCustomer,
	logger *slog.Logger,
) {
	jobs := []struct {
		name string
		url  string
		fn   func(args []byte) error
	}{
		{
			name: "consumer-notify-user",
			url:  os.Getenv("AWS_SQS_QUEUE_URL"),
			fn: func(args []byte) error {
				var input *usecases.NotifyCustomerInput
				err := json.Unmarshal(args, &input)
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				err = notifyUser.Execute(input)
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				return nil
			},
		},
	}
	for _, job := range jobs {
		go queue.Consume(job.url, job.fn)
	}
}
