package domain

import (
	"app/internal/application/exceptions"
	"net/mail"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func (u *User) Validate() error {
	if isInvalidEmail(u.Email) {
		return exceptions.NewDomainException(exceptions.INVALID_EMAIL)
	}
	return nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
