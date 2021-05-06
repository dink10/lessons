package main

import (
	"context"
	"github.com/dink10/lessons/interface2/db"
)

type MySql struct {
}

func (MySql) Create(ctx context.Context, item *db.Item) (*db.Item, error) {
	return nil, nil
}

func (MySql) Update(ctx context.Context, item *db.Item) error {
	return nil
}

func (MySql) Delete(ctx context.Context, id int) error {
	return nil
}

func (MySql) Read(ctx context.Context) []db.Item {
	return nil
}
