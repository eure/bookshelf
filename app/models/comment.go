package models

import "time"

// Comment is
type Comment struct {
	ID      int       `json:"id" xorm:"id"`
	BookID  int       `json:"book_id" xorm:"book_id"`
	Body    string    `json:"body" xorm:"body"`
	UserID  int       `json:"user_id" xorm:"user_id"`
	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}

// NewComment ...
func NewComment(bookID int, body string) Comment {
	return Comment{
		BookID: bookID,
		Body:   body,
	}
}

// CommentRepository is
type CommentRepository struct {
}

// NewCommentRepository ...
func NewCommentRepository() CommentRepository {
	return CommentRepository{}
}

// GetByID ...
func (r CommentRepository) GetByID(id int) (*Comment, error) {
	comment := Comment{ID: id}
	has, err := engine.Get(&comment)
	if err != nil {
		return nil, err
	}
	if has {
		return &comment, nil
	}
	return nil, nil
}

// Insert ...
func (r CommentRepository) Insert(bookID int, body string, userID int) (*Comment, error) {
	comment := NewComment(bookID, body)
	comment.UserID = userID
	if _, err := engine.Nullable("user_id").InsertOne(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

// FindByBooks ...
func (r CommentRepository) FindByBooks(books []Book) (map[int][]Comment, error) {
	bookIDs := make([]int, len(books))
	for i, book := range books {
		bookIDs[i] = book.ID
	}
	var comments []Comment
	if err := engine.In("book_id", bookIDs).Find(&comments); err != nil {
		return nil, err
	}
	res := make(map[int][]Comment)
	for _, comment := range comments {
		res[comment.BookID] = append(res[comment.BookID], comment)
	}
	return res, nil
}
