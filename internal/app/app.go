package app

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/clipclock08/news-crud/internal/config"
	"github.com/clipclock08/news-crud/internal/services/article"
	repo "github.com/clipclock08/news-crud/internal/services/article/mongo"
)

type Options func(*App) error

type App struct {
	Cfg            config.Config
	ArticleService *article.Service
}

func New(opts ...Options) (*App, error) {
	as := &App{}
	for _, opt := range opts {
		if err := opt(as); err != nil {
			return nil, err
		}
	}

	return as, nil
}

func WithMongoRepo(db *mongo.Client) Options {
	return func(app *App) error {
		as, err := article.New(article.WithRepo(repo.NewMongoRepo(db, app.Cfg.DB.Name)))
		if err != nil {
			return err
		}

		app.ArticleService = as
		return nil
	}
}
