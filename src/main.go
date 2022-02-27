package main

import (
	"log"
	"shortner/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	sc := controller.NewShortnerController()

	router.LoadHTMLGlob("./views/*")
	router.GET("/", sc.Index)
	router.GET("/short", sc.Create)
	router.GET("/list", sc.List)

	router.POST("/api/urls", sc.ShortUrl)
	router.GET("/api/urls", sc.GetUrls)

	router.GET("/r/:id", sc.Redirect)

	if err := router.Run("0.0.0.0:80"); err != nil {
		log.Fatalf("Failed to start gin server, got error %+v", err)
	}
}
