package model

type UserRepository interface {
	FindById()
	FindAll() []string //@TODO remove after tests
	Create()
	Update()
}
