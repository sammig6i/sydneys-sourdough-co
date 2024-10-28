package repository

import (
	"context"
	"fmt"

	"github.com/pgvector/pgvector-go"
	"github.com/sammig6i/sydneys-sourdough-co/database"
	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type searchRepository struct {
	db database.Database
}

func NewSearchRepository(db database.Database) domain.SearchRepository {
	return &searchRepository{db: db}
}

func (s *searchRepository) Search(ctx context.Context, query string) ([]*domain.SearchResult, error) {
	queryEmbedding := pgvector.NewVector(getEmbedding(query))

	rows, err := s.db.Query(ctx, `
		SELECT 
			entity_type,
			id,
			name,
			context,
			1 - (embedding <=> $1) as score
		FROM unified_search
		ORDER BY embedding <=> $1
	`, queryEmbedding)

	if err != nil {
		return nil, fmt.Errorf("error performing search: %w", err)
	}
	defer rows.Close()

	var results []*domain.SearchResult
	for rows.Next() {
		var result domain.SearchResult
		err := rows.Scan(
			&result.EntityType,
			&result.ID,
			&result.Name,
			&result.Context,
			&result.Score,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning search result: %w", err)
		}
		results = append(results, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating search results: %w", err)
	}

	return results, nil
}
