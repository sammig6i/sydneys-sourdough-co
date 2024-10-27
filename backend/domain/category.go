package domain

import "context"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Embedding []float32 `json:"-"`
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]*Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*Category, error)
	GetByName(c context.Context, name string) (*Category, error)
}

type CategoryUsecase interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]*Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*Category, error)
	GetByName(c context.Context, name string) (*Category, error)
}
