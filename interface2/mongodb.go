package main

import (
	"context"

	"github.com/dink10/lessons/interface2/db"
)

type Mongo struct {
}

func (Mongo) Create(ctx context.Context, item *db.Item) (*db.Item, error) {
	return nil, nil
}

func (Mongo) Update(ctx context.Context, item *db.Item) error {
	return nil
}

func (Mongo) Delete(ctx context.Context, id int) error {
	return nil
}

func (Mongo) Read(ctx context.Context) []db.Item {
	return nil
}
