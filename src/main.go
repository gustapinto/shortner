package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/short", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.tmpl", gin.H{})
	})
	router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.tmpl", gin.H{})
	})

	if err := router.Run("0.0.0.0:80"); err != nil {
		log.Fatalf("Failed to start gin server, got error %+v", err)
	}
}
