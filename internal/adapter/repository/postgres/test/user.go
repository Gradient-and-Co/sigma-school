package test

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/guregu/null"
)

type UserBuilder struct {
	user domain.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: domain.User{
			ID:        domain.NewID(),
			Name:      "John",
			Surname:   "Doe",
			Phone:     null.StringFrom("+123456789"),
			City:      null.StringFrom("Berlin"),
			AvatarUrl: null.StringFrom("https://example.com/avatar"),
			Email:     "john.doe@example.com",
			Password:  "password",
		},
	}
}

func (b *UserBuilder) WithID(id domain.ID) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) Build() domain.User {
	return b.user
}
