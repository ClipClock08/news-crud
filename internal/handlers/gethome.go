package handlers

import (
	"github.com/clipclock08/news-crud/internal/models"
	"github.com/clipclock08/news-crud/internal/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := views.Index([]*models.Article{}).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
