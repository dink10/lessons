package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dink10/lessons/server/storage"
	"net/http"

	"github.com/go-chi/chi"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	res := map[string]interface{}{}

	storage.Storage.Range(func(key, value interface{}) bool {
		res[fmt.Sprint(key)] = value

		return true
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(res)
	_ = err
	_, _ = w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	val, ok := storage.Storage.Load(id)
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		_, _ = w.Write(nil)
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(val)
	_, _ = w.Write(data)
}
