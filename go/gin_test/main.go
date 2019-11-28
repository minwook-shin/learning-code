package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/default", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	engine.Use(gin.Logger())

	engine.Use(gin.Recovery())

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	engine.GET("/param/:test", func(c *gin.Context) {
		val := c.Param("test")
		c.String(http.StatusOK, val)
	})

	engine.GET("/param/:test/*action", func(c *gin.Context) {
		val := c.Param("test")
		action := c.Param("action")
		message := val + " " + action

		fmt.Println(c.FullPath() == "/param/:test/*action")

		c.String(http.StatusOK, message)
	})

	engine.GET("/query", func(c *gin.Context) {
		DefaultQuery := c.DefaultQuery("DefaultQuery", "true")
		query := c.Query("q")

		c.String(http.StatusOK, DefaultQuery, query)
	})

	engine.POST("/post", func(c *gin.Context) {
		postForm := c.PostForm("PostForm")
		defaultPostForm := c.DefaultPostForm("DefaultPostForm", "true")

		c.JSON(200, gin.H{
			"PostForm":        postForm,
			"DefaultPostForm": defaultPostForm,
		})
	})

	engine.POST("/post_query", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("Default", "0")
		c.String(http.StatusOK, id, page)
	})

	engine.POST("/post_map", func(c *gin.Context) {
		queryMap := c.QueryMap("map")
		postFormMap := c.PostFormMap("PostFormMap")

		fmt.Println(queryMap, postFormMap)
	})

	engine.POST("/file", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(http.StatusOK, fmt.Sprintf("'%s'", file.Filename))
	})

	engine.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})

	engine.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	engine.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		engine.HandleContext(c)
	})
	engine.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	engine.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	engine.Run(":8080")
}
