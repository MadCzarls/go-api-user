package service

type VariableGetter interface {
	GetEnvString(key string) *string
	GetEnvInt(key string) *int
}
