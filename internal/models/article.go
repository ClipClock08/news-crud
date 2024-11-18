package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type UpdateArticleParams struct {
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

func NewArticle(title, content string) (*Article, error) {
	if title == "" {
		return nil, errors.New("article title can not be empty")
	}

	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Article{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
