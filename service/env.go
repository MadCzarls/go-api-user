package service

import (
	"fmt"
	"github.com/joho/godotenv"
	"strconv"
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

func (em *envManager) GetEnvString(key string) *string {
	if _, present := em.envs[key]; !present {
		panic(fmt.Sprintf("env variable for key '%s' does not exist", key))
	}

	value := em.envs[key]

	return &value
}

func (em *envManager) GetEnvInt(key string) *int {
	if _, present := em.envs[key]; !present {
		panic(fmt.Sprintf("env variable for key '%s' does not exist", key))
	}

	value := em.envs[key]
	valueInt, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return &valueInt
}
