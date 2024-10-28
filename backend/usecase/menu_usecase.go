package usecase

import (
	"context"
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type menuItemUsecase struct {
	menuItemRepo   domain.MenuItemRepository
	contextTimeout time.Duration
}

func NewMenuUsecase(menuItemRepo domain.MenuItemRepository, timeout time.Duration) domain.MenuItemUsecase {
	return &menuItemUsecase{
		menuItemRepo:   menuItemRepo,
		contextTimeout: timeout,
	}
}

func (mu *menuItemUsecase) Create(ctx context.Context, menuItem *domain.MenuItem) error {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.Create(ctx, menuItem)
}

func (mu *menuItemUsecase) Fetch(ctx context.Context) ([]*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.Fetch(ctx)
}

func (mu *menuItemUsecase) Update(ctx context.Context, menuItem *domain.MenuItem) error {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.Update(ctx, menuItem)
}

func (mu *menuItemUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.Delete(ctx, id)
}

func (mu *menuItemUsecase) GetByID(ctx context.Context, id int) (*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.GetByID(ctx, id)
}

func (mu *menuItemUsecase) GetByCategory(ctx context.Context, categoryID int) ([]*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.GetByCategory(ctx, categoryID)
}

func (mu *menuItemUsecase) GetByPriceRange(ctx context.Context, minPrice, maxPrice float64) ([]*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepo.GetByPriceRange(ctx, minPrice, maxPrice)
}
