package usecase

import (
	"context"
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type searchUsecase struct {
	searchRepo     domain.SearchRepository
	contextTimeout time.Duration
}

func NewSearchUsecase(searchRepo domain.SearchRepository, timeout time.Duration) domain.SearchUsecase {
	return &searchUsecase{
		searchRepo:     searchRepo,
		contextTimeout: timeout,
	}
}

func (su *searchUsecase) Search(ctx context.Context, query string) ([]*domain.SearchResult, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.searchRepo.Search(ctx, query)
}
