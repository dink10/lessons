package db

import "context"

type Repository interface {
	Create(ctx context.Context, item *Item) (*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id int) error
	Read(ctx context.Context) []Item
}
