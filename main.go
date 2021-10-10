package main

import (
	"fmt"
	"net/http"

	_ "github.com/practical-coder/booknotes/db"
	"github.com/practical-coder/booknotes/handlers/notes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := SetEngine()
	engine.Run()
}

func SetEngine() *gin.Engine {
	e := gin.Default()
	// Notes
	e.POST("/notes", notes.Create)
	e.GET("/notes", notes.List)
	e.GET("/notes/:uuid", notes.Show)
	e.PUT("/notes/:uuid", notes.Update)
	e.DELETE("/notes/:uuid", notes.Delete)
	e.GET("/notes/search", notes.Search)

	// hello
	e.GET("/:directory", helloHandler)

	return e
}

func helloHandler(c *gin.Context) {
	d := c.Params.ByName("directory")
	msg := fmt.Sprintf("It works! Directory: %s", d)
	body := gin.H{
		"message": msg,
	}
	c.JSON(http.StatusOK, body)
}
