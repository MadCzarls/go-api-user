package model

type UserRepository interface {
	FindById(id string) (*User, error)
	FindAll() ([]User, error)
	Create(*User) error
	Update(id string, user *User) error
}

type Idier interface {
	SetId(id string)
	GetId() string
}
