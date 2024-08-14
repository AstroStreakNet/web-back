package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Password    string
	DisplayName string
	FirstName   *string
	LastName    *string
	Role        string
	Images      []Image // Necessary for HAS MANY relationship in GORM
}

// Builder

type UserBuilder struct {
	email       string
	password    string
	displayName string
	firstName   *string
	lastName    *string
	role        string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (builder *UserBuilder) WithEmail(email string) *UserBuilder {
	builder.email = email
	return builder
}

func (builder *UserBuilder) WithPassword(password string) *UserBuilder {
	builder.password = password
	return builder
}

func (builder *UserBuilder) WithDisplayName(displayName string) *UserBuilder {
	builder.displayName = displayName
	return builder
}

func (builder *UserBuilder) WithFirstName(firstName string) *UserBuilder {
	builder.firstName = &firstName
	return builder
}

func (builder *UserBuilder) WithLastName(lastName string) *UserBuilder {
	builder.lastName = &lastName
	return builder
}

func (builder *UserBuilder) WithRole(role string) *UserBuilder {
	builder.role = role
	return builder
}

func (builder *UserBuilder) Build() *User {
	return &User{
		Email:       builder.email,
		Password:    builder.password,
		DisplayName: builder.displayName,
		FirstName:   builder.firstName,
		LastName:    builder.lastName,
		Role:        builder.role,
	}
}
