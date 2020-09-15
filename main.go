package main

import (
	"gin_test/controller"
	"gin_test/middleware"
	"io"
	"log"
	"net/http"
	"os"

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

	here_document := `<!DOCTYPE html>
	<html lang="ja">
	<head>
	<meta charset="uft-8">
	<title>file upload</title>
	</head>
	<body>
		<form action="upload" method="post" enctype="multipart/form-data">
			<p>
			<input type="file" name="image">
			</p>
			<p>
			<input type="submit" value="upload">
			</p>
			</form>
	</body>
	</html>`

	engine.GET("/upload", func(c *gin.Context) {
		c.Writer.WriteString(here_document)
	})
	// engine.GET("/upload", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "fileupload.html", gin.H{
	// 		"title": "file upload",
	// 	})
	// })

	engine.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		fileName := header.Filename
		dir, _ := os.Getwd()
		out, err := os.Create(dir + "\\images\\" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	engine.Use(middleware.RecordUaAndTime)

	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}

	engine.Run(":3000")
}
