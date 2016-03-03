package main

import (
	"bookshelf/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

// main ...
func main() {

	router := gin.Default()
	router.GET("/:id", func(c *gin.Context) {
		// Pramを処理する
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"status": "id shoud be bigger than 0"})
			return
		}
	})
	// データを処理する
	ctrl := controllers.NewUser()
	result := ctrl.Get(id)
	if result == nil || reflect.ValueOf(result).IsNil() {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, result)
	router.Run(":8080")
}
