package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/eure/bookshelf/app/models"
)

// GetBooks ...
func GetBooks(c *gin.Context) {
	repo := models.NewBookRepository()
	books, err := repo.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	commentRepo := models.NewCommentRepository()
	comments, err := commentRepo.FindByBooks(books)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	type response struct {
		Book     models.Book      `json:"book"`
		Comments []models.Comment `json:"comments"`
	}
	responses := make([]response, len(books))
	for i, book := range books {
		responses[i] = response{
			Book:     book,
			Comments: comments[book.ID],
		}
	}
	c.JSON(http.StatusOK, responses)
}

// PostComment ...
func PostComment(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	repo := models.NewBookRepository()
	book, err := repo.GetByID(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	body := c.PostForm("body")
	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("body is not full"))
		return
	}

	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	if userID != 0 {
		userRepo := models.NewUserRepository()
		user, err := userRepo.GetByID(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if user == nil {
			c.JSON(http.StatusBadRequest, errors.New("user is not found"))
			return
		}
	}

	commentRepo := models.NewCommentRepository()
	comment, err := commentRepo.Insert(book.ID, body, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, comment)
}
