package model

type UserRepository interface {
	FindById(id string) (*User, error)
	FindAll() ([]User, error)
	Create(*User) error
	Update()
}

type Idier interface {
	SetId(id string)
}
