package adapters

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Sqs struct {
	client *sqs.SQS
}

func (s *Sqs) Connect() error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: *aws.String(os.Getenv("AWS_PROFILE")),
		Config: aws.Config{
			Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
			Region:   aws.String(os.Getenv("AWS_REGION")),
		},
	})
	if err != nil {
		return err
	}
	s.client = sqs.New(sess)
	return nil
}

func (s *Sqs) Publish(queueName string, data []byte) error {
	_, err := s.client.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(data)),
		QueueUrl:    aws.String(queueName),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqs) Consume(queueName string, callback func(args []byte)) error {
	for {
		ouptut, err := s.client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(queueName),
		})
		if err != nil {
			return err
		}
		for _, msg := range ouptut.Messages {
			body := []byte(*msg.Body)
			callback(body)
			s.client.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueName),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}
	}
}
