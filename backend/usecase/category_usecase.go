package usecase

import (
	"context"
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type categoryUsecase struct {
	categoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo:   categoryRepo,
		contextTimeout: timeout,
	}
}

func (cu *categoryUsecase) Create(ctx context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepo.Create(ctx, category)
}

func (cu *categoryUsecase) Fetch(ctx context.Context) ([]*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepo.Fetch(ctx)
}

func (cu *categoryUsecase) Update(ctx context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepo.Update(ctx, category)
}

func (cu *categoryUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepo.Delete(ctx, id)
}

func (cu *categoryUsecase) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepo.GetByID(ctx, id)
}
