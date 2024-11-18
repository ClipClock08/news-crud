package article

import (
	"context"
	"errors"
	"time"

	"github.com/clipclock08/news-crud/internal/models"
	"github.com/clipclock08/news-crud/internal/services/article/repo"
)

type Service struct {
	repo repo.Article
}

type Option func(*Service) error

func New(opts ...Option) (*Service, error) {
	s := &Service{}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithRepo(r repo.Article) Option {
	return func(s *Service) error {
		if r == nil {
			return errors.New("nil repo")
		}
		s.repo = r
		return nil
	}
}

func (s *Service) CreateArticle(ctx context.Context, title, body string) error {
	return nil
}

func (s *Service) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	return nil, nil
}

func (s *Service) GetArticles(ctx context.Context, offset, limit int) ([]*models.Article, error) {
	return nil, nil
}

func (s *Service) UpdateArticle(ctx context.Context, params models.UpdateArticleParams, updatedAt time.Time) error {
	return nil
}

func (s *Service) DeleteArticle(ctx context.Context, id string) error {
	return nil
}
