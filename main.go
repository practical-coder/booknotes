package main

import (
	"fmt"
	_ "github.com/practical-coder/booknotes/db"
	"github.com/practical-coder/booknotes/handlers/notes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Notes
	router.POST("/notes", notes.Create)
	router.GET("/notes", notes.List)
	router.GET("/notes/:uuid", notes.Show)
	router.PUT("/notes/:uuid", notes.Update)
	router.DELETE("/notes/:uuid", notes.Delete)
	router.GET("/notes/search", notes.Search)

	// hello
	router.GET("/:directory", helloHandler)

	router.Run()
}

func helloHandler(c *gin.Context) {
	d := c.Params.ByName("directory")
	msg := fmt.Sprintf("It works! Directory: %s", d)
	body := gin.H{
		"message": msg,
	}
	c.JSON(http.StatusOK, body)
}
