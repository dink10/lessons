package handler

import (
	"github.com/dink10/lessons/server/storage"
	"github.com/go-chi/chi"
	"net/http"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	storage.Storage.Delete(id)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
