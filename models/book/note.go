package book

import (
	"context"
	"github.com/practical-coder/booknotes/db"
	"github.com/practical-coder/booknotes/zlg"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UUID        string             `json:"uuid"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Content     string             `json:"content"`
	Authors     []string           `json:"authors"`
	Topics      []string           `json:"topics"`
	Tags        []string           `json:"tags"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	PublishedAt time.Time          `json:"published_at"`
}

func notes() *mongo.Collection {
	return db.Client.Database("booknotes").Collection("notes")
}

func (n *Note) AssignID() {
	n.ID = primitive.NewObjectID()
}

func (n *Note) AssignUUID() {
	n.UUID = uuid.New().String()
}

func (n *Note) InitialDates() {
	n.touchCreatedAt()
	n.UpdatedAt = n.CreatedAt
}

func (n *Note) touchCreatedAt() {
	date := time.Now()
	n.CreatedAt = date

}
func (n *Note) TouchUpdatedAt() {
	date := time.Now()
	n.UpdatedAt = date
}

func CreateNote(n *Note) (*Note, error) {
	n.AssignID()
	n.AssignUUID()
	n.InitialDates()

	result, err := notes().InsertOne(context.Background(), n)
	if err != nil {
		zlg.Logger.Error().Err(err).Str("action", "create_note").Msg("CreateNote Error")
		return n, err
	}
	zlg.Logger.Info().Str("action", "create_note").Interface("result", result).Msg("CreateNote Result")

	return n, nil
}

func (n *Note) Merge(reqNote Note) {
	if reqNote.Title != "" {
		n.Title = reqNote.Title
	}

	if reqNote.Description != "" {
		n.Description = reqNote.Description
	}

	if reqNote.Authors != nil {
		n.Authors = reqNote.Authors
	}

	if reqNote.Content != "" {
		n.Content = reqNote.Content
	}

	if reqNote.Topics != nil {
		n.Topics = reqNote.Topics
	}

	if reqNote.Tags != nil {
		n.Tags = reqNote.Tags
	}
}

func FindNotes(ctx context.Context, filter bson.D) ([]Note, error) {
	results := make([]Note, 0)
	cursor, err := notes().Find(ctx, filter)
	if err != nil {
		zlg.Logger.Error().Err(err).Str("action", "find_notes").Msg("FindNotes Error")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var note Note
		err := cursor.Decode(&note)
		if err != nil {
			zlg.Logger.Error().Err(err).Str("action", "find_notes").Msg("FindNotes Cursor Decode Error")
		}
		results = append(results, note)
	}
	return results, nil
}

func FindNote(ctx context.Context, filter bson.D) (*Note, error) {
	note := new(Note)
	err := notes().FindOne(ctx, filter).Decode(note)
	if err != nil {
		zlg.Logger.Error().Err(err).Str("action", "find_note").Msg("FindNote FindOne/Decode Error")
		return nil, err
	}
	return note, nil
}

func DeleteNote(ctx context.Context, filter bson.D) (*mongo.DeleteResult, error) {
	result, err := notes().DeleteOne(ctx, filter)
	if err != nil {
		zlg.Logger.Error().
			Err(err).
			Str("action", "delete_note").
			Msg("DeleteNote DeleteOne Error")
		return nil, err
	}
	zlg.Logger.Info().
		Str("action", "delete_note").
		Interface("delete_filter", filter).
		Interface("delete_result", result).
		Msg("DeleteNote Result")

	return result, nil
}

func UpdateNote(ctx context.Context, filter bson.D, note *Note) (*mongo.UpdateResult, error) {
	result, err := notes().UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: note}})
	if err != nil {
		zlg.Logger.Error().
			Err(err).
			Str("action", "update_note").
			Msg("UpdateNote UpdateOne Error")
		return nil, err
	}
	zlg.Logger.Info().
		Str("action", "update_note").
		Interface("update_filter", filter).
		Interface("update_object", note).
		Interface("update_result", result).
		Msg("UpdateNote Result")

	return result, nil
}
