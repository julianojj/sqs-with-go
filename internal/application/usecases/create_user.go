package usecases

import (
	"app/internal/application/domain"
	"app/internal/ports"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

type CreateUser struct {
	queue  ports.Queue
	logger *slog.Logger
}

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCreateUser(
	queue ports.Queue,
	logger *slog.Logger,
) *CreateUser {
	return &CreateUser{
		queue,
		logger,
	}
}

func (c *CreateUser) Execute(input *CreateUserInput) error {
	user := &domain.User{
		Id:    uuid.NewString(),
		Name:  input.Name,
		Email: input.Email,
	}
	err := user.Validate()
	if err != nil {
		c.logger.Error(
			"Error to create user",
			"reason", err.Error(),
			"call", "create_user.execute",
		)
		return err
	}
	c.logger.Info(
		"User created sucessfully",
		"log", user,
		"call", "create_user.execute",
	)
	data := map[string]any{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = c.queue.Publish(
		os.Getenv("AWS_SQS_QUEUE_URL"),
		b,
	)
	if err != nil {
		return err
	}
	return nil
}
