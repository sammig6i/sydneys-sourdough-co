package domain

import "context"

type SearchResult struct {
	EntityType string  `json:"entity_type"`
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Context    string  `json:"context,omitempty"`
	Score      float64 `json:"score"`
}

type SearchRepository interface {
	Search(ctx context.Context, query string) ([]*SearchResult, error)
}

type SearchUsecase interface {
	Search(ctx context.Context, query string) ([]*SearchResult, error)
}
