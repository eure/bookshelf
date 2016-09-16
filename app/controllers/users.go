package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eure/bookshelf/app/models"
)

// PostUser ...
func PostUser(c *gin.Context) {
	name := c.PostForm("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("name is not full"))
		return
	}
	repo := models.NewUserRepository()
	user, err := repo.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if user != nil {
		c.JSON(http.StatusBadRequest, errors.New("your name already exits"))
		return
	}
	password := c.PostForm("password")
	if len(password) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("password is not full"))
		return
	}
	rePassword := c.PostForm("re_password")
	if password != rePassword {
		c.JSON(http.StatusBadRequest, errors.New("password is not equal"))
		return
	}

	user, err = repo.Insert(name, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// PostUserLogin ...
func PostUserLogin(c *gin.Context) {
	name := c.PostForm("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("name is not full"))
		return
	}
	repo := models.NewUserRepository()
	user, err := repo.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, errors.New("user is not found"))
		return
	}

	password := c.PostForm("password")
	if len(password) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("password is not full"))
		return
	}
	if err := user.ValidatePassword(password); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
