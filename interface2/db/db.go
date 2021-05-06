package db

import (
	"context"
)

type DB struct {
	DB Repository
}

func (db *DB) Save(ctx context.Context, item Item) error {
	_, err := db.DB.Create(ctx, &item)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Delete(ctx context.Context, item Item) error {
	return db.DB.Delete(ctx, item.id)
}
