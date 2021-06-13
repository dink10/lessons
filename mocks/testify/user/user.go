package user

import (
	"github.com/dink10/lessons/mocks/testify/doer"
)

type User struct {
	Doer doer.Doer
}

func (u *User) Use() string {
	return u.Doer.Do(1, "work1")
}
