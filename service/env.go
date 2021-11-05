package service

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

type envManager struct {
	envs map[string]string
}

func NewEnvManager() (*envManager, error) {
	envs, err := godotenv.Read()

	if err != nil {
		return nil, err
	}

	return &envManager{envs: envs}, nil
}

func (em *envManager) GetVariable(key string) (*string, error) {
	value := em.envs[key]

	if value == "" {
		return nil, errors.New(fmt.Sprintf("env variable for key '%s' does not exist", key))
	}

	return &value, nil
}
