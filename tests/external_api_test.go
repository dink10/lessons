package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExternalAPI(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"id":1, "name":"TestName", "rating": 10}`)
	}))

	defer ts.Close()

	user := User{Id: 2, Name: "empty"}

	fmt.Println(ts.URL)

	err := ApiCaller(&user, ts.URL)
	if err != nil {
		t.Error(err)
	}
}
