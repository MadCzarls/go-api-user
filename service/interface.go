package service

type VariableGetter interface {
	GetVariable(key string) (*string, error)
}
