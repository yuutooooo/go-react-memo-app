package model

import (
	"time"
)

type User struct {
	id        string
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

// NewUser Userエンティティを作成
func NewUser(name, email, password string) *User {
	return &User{
		name:      name,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// Getters
func (u *User) ID() string           { return u.id }
func (u *User) Name() string         { return u.name }
func (u *User) Email() string        { return u.email }
func (u *User) Password() string     { return u.password }
func (u *User) CreatedAt() time.Time { return u.createdAt }
func (u *User) UpdatedAt() time.Time { return u.updatedAt }

// Setters
func (u *User) SetName(name string) {
	u.name = name
	u.updatedAt = time.Now()
}
func (u *User) SetEmail(email string) {
	u.email = email
	u.updatedAt = time.Now()
}
func (u *User) SetPassword(password string) {
	u.password = password
	u.updatedAt = time.Now()
}

func (u *User) SetID(id string) {
	u.id = id
	u.updatedAt = time.Now()
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
