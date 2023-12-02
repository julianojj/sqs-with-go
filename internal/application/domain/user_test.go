package domain

import (
	"app/internal/application/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "new user",
			fn: func(t *testing.T) {
				user := &User{
					Id:    "1",
					Name:  "John Doe",
					Email: "john.doe@test.com",
				}
				assert.Equal(t, "1", user.Id)
				assert.Equal(t, "John Doe", user.Name)
				assert.Equal(t, "john.doe@test.com", user.Email)
			},
		},
		{
			name: "return exception if invalid email",
			fn: func(t *testing.T) {
				user := &User{
					Id:    "1",
					Name:  "John Doe",
					Email: "john.doetest.com",
				}
				err := user.Validate()
				assert.EqualError(t, err, exceptions.INVALID_EMAIL)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
