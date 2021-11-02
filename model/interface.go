package model

type UserRepository interface {
	FindById()
	FindAll() ([]User, error)
	Create(User) error
	Update()
}

type Idier interface {
	SetId(id string)
}