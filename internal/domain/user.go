package domain

import (
	"regexp"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func NewUser(id int, name, email string) (*User, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}

	if name == "" || len(name) < 5 {
		return nil, ErrInvalidName
	}

	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmail
	}

	now := time.Now()

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
