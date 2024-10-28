package domain

import "context"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Embedding []float32 `json:"-"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	Fetch(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*Category, error)
}

type CategoryUsecase interface {
	Create(ctx context.Context, category *Category) error
	Fetch(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*Category, error)
}
