package notes

import (
	"context"
	"net/http"
	"time"

	"github.com/practical-coder/booknotes/models/book"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func List(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	notes, err := book.FindNotes(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, notes)
}

func Create(c *gin.Context) {
	note := new(book.Note)
	err := c.ShouldBindJSON(note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	note, err = book.CreateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, note)
}

func Show(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	uuid := c.Params.ByName("uuid")
	filter := bson.D{{Key: "uuid", Value: uuid}}

	note, err := book.FindNote(ctx, filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, note)
}

func Update(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	filter := bson.D{{Key: "uuid", Value: uuid}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reqNote book.Note
	err := c.ShouldBindJSON(&reqNote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	note, err := book.FindNote(ctx, filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	note.Merge(reqNote)
	note.TouchUpdatedAt()

	result, err := book.UpdateNote(ctx, filter, note)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func Delete(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	filter := bson.D{{Key: "uuid", Value: uuid}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := book.DeleteNote(ctx, filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func Search(c *gin.Context) {
	tag := c.Query("tag")
	filter := bson.D{{Key: "tag", Value: tag}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notes, err := book.FindNotes(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, notes)
}
