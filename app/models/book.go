package models

import "time"

// Book is
type Book struct {
	ID      int       `json:"id" xorm:"id"`
	Name    string    `json:"name" xorm:"name"`
	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}

// NewBook ...
func NewBook(name string) Book {
	return Book{
		Name: name,
	}
}

// BookRepository is
type BookRepository struct {
}

// NewBookRepository ...
func NewBookRepository() BookRepository {
	return BookRepository{}
}

// GetByID ...
func (r BookRepository) GetByID(id int) (*Book, error) {
	book := Book{ID: id}
	has, err := engine.Get(&book)
	if err != nil {
		return nil, err
	}
	if has {
		return &book, nil
	}
	return nil, nil
}

// Find ...
func (r BookRepository) Find() ([]Book, error) {
	var books []Book
	if err := engine.Find(&books); err != nil {
		return nil, err
	}
	return books, nil
}
