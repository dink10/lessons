package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiCaller(user *User, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_ = json.NewDecoder(resp.Body).Decode(user)

	fmt.Println(user)

	return nil
}
