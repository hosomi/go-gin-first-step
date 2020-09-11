package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world!",
			"User-Agent": ua,
		})
	})

	engine.LoadHTMLGlob("templates/*")
	engine.GET("/templates", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello world!",
			"ua":      ua,
		})
	})

	engine.Static("/static", "./static")

	engine.Run(":3000")
}
