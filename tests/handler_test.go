package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost/user?id=42", nil)
	w := httptest.NewRecorder()
	userHandler(w, r)

	user := User{}
	_ = json.Unmarshal(w.Body.Bytes(), &user)
	if user.Id != 42 {
		t.Errorf("Invalid user id %d, expected: %d", user.Id, 42)
	}
}
