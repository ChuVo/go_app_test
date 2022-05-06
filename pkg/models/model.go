package models

import (
	"errors"
)

var ErrNoUser = errors.New("models:  Пользователь не найден")

type User struct {
	ID        int
	UserName  string
	UserEmail string
	Admin     bool
	Delete    bool
}

type UserList struct {
	User *[]User
}
