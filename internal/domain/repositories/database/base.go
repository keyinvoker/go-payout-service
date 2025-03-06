package repositories

import "context"

type BaseDatabaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id int) (*T, error)
	GetAll(ctx context.Context, filters map[string]any) ([]*T, error)
}
