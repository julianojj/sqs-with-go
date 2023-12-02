package usecases

import (
	"log/slog"
)

type NotifyCustomer struct {
	logger *slog.Logger
}

type NotifyCustomerInput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewNotifyCustomer(logger *slog.Logger) *NotifyCustomer {
	return &NotifyCustomer{
		logger,
	}
}

func (nc *NotifyCustomer) Execute(input *NotifyCustomerInput) error {
	nc.logger.Info(
		"Sending message to user",
		"log", input,
	)
	return nil
}
