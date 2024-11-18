package repo

import (
	"context"
	"time"

	"github.com/clipclock08/news-crud/internal/models"
)

type Article interface {
	CreateArticle(ctx context.Context, title, body string) error
	GetArticle(ctx context.Context, id string) (models.Article, error)
	GetArticles(ctx context.Context, offset, limit int) ([]models.Article, error)
	UpdateArticle(ctx context.Context, params models.UpdateArticleParams, updatedAt time.Time) error
	DeleteArticle(ctx context.Context, id string) error
}
