package entity

type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

type Users []*User
