package main

import (
	"github.com/eure/kamimai"
	"github.com/eure/kamimai/core"
	_ "github.com/eure/kamimai/driver"
	"github.com/gin-gonic/gin"

	"github.com/eure/bookshelf/app/routes"
)

func syncDatabase() {
	conf, err := core.NewConfig("database/mysql")
	if err != nil {
		panic(err)
	}
	conf.WithEnv("development")
	kamimai.Sync(conf)
}

// main ...
func main() {
	if gin.IsDebugging() {
		syncDatabase()
	}
	routes.Handler().Run(":8080")
}
