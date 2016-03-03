package models

import (
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

// main ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "fuga:pass@/bookshelf")
	if err != nil {
		panic(err)
	}
}

// UserRepository is
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

// NewUser ...
func NewUser(id int, username string) User {
	return User{
		ID:       id,
		Username: username,
	}
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}

	return nil
}
