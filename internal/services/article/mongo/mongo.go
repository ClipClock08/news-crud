package mongo

import (
	"context"
	"github.com/clipclock08/news-crud/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepo struct {
	db *mongo.Database
}

func NewMongoRepo(client *mongo.Client, daName string) *MongoRepo {
	return &MongoRepo{
		db: client.Database(daName),
	}
}

func (m MongoRepo) CreateArticle(ctx context.Context, title, body string) error {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) GetArticle(ctx context.Context, id string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) GetArticles(ctx context.Context, offset, limit int) (posts []models.Article, err error) {
	opts := options.Find()
	skip := (offset - 1) * limit
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))

	cursor, err := m.db.Collection("articles").Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post models.Article
		if err = cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (m MongoRepo) UpdateArticle(ctx context.Context, params models.UpdateArticleParams, updatedAt time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) DeleteArticle(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
