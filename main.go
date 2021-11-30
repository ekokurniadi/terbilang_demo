package main

import (
	"github.com/ekokurniadi/terbilang_demo/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/**/*")
	handlerApi := handler.NewRouterHandler()

	api := router.Group("api/v1")
	api.POST("/proses", handlerApi.Process)
	router.GET("/", handlerApi.Index)
	router.Run()
}
