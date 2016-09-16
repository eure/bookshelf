package models

import (
	"errors"
	"time"
)

// User is
type User struct {
	ID       int       `json:"id" xorm:"id"`
	Name     string    `json:"name" xorm:"name"`
	Password string    `json:"-" xorm:"password"`
	Created  time.Time `json:"created" xorm:"created"`
	Updated  time.Time `json:"updated" xorm:"updated"`
}

// NewUser ...
func NewUser(name, password string) User {
	return User{
		Name:     name,
		Password: password,
	}
}

func (u User) ValidatePassword(password string) error {
	if u.Password != password {
		return errors.New("password is not correct")
	}
	return nil
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (r UserRepository) GetByID(id int) (*User, error) {
	user := User{ID: id}
	has, err := engine.Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}
	return nil, nil
}

// GetByName ...
func (r UserRepository) GetByName(name string) (*User, error) {
	user := User{}
	has, err := engine.Where("name = ?", name).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}
	return nil, nil
}

// Insert ...
func (r UserRepository) Insert(name, password string) (*User, error) {
	user := NewUser(name, password)
	if _, err := engine.InsertOne(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
