package handler

import (
	"encoding/json"
	"github.com/dink10/lessons/server/storage"
	"io"
	"net/http"

	"github.com/dink10/lessons/server/model"
	"github.com/google/uuid"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var product model.Product
	if err := json.Unmarshal(b, &product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	u, err := uuid.NewUUID()
	if err := json.Unmarshal(b, &product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	product.ID = u.String()

	storage.Storage.Store(product.ID, product)

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(product)
	_, _ = w.Write(data)
}
