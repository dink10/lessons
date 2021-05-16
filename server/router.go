package main

import (
	"github.com/dink10/lessons/server/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	// CRD
	r.Get("/api/v1/products", handler.GetProducts)
	r.Get("/api/v1/products/{id}", handler.GetProduct)
	r.Post("/api/v1/product", handler.AddProduct)
	r.Delete("/api/v1/product/{id}", handler.DeleteProducts)

	return r
}
