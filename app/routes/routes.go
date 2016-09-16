package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/eure/bookshelf/app/controllers"
)

func Handler() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/books", controllers.GetBooks)
	router.POST("/books/:id/comment", controllers.PostComment)
	router.POST("/users", controllers.PostUser)
	router.POST("/users/login", controllers.PostUserLogin)
	return
}
