package controllers

import (
	"bookshelf/models"
)

// User is
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Get ...
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	return user
}
