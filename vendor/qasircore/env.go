package qasircore

import (
	"fmt"
	"os"
	"strings"
)

type Environtment struct {
	configs map[string]interface{}
}

func (e *Environtment) Set(key, value string) {
	e.configs[key] = value
	os.Setenv(key, value)
}

func (e *Environtment) Get(key string) string {
	return os.Getenv(key)
}

func (e *Environtment) Load() {
	e.loadFile(false)
}

func (e *Environtment) Overload() {
	e.loadFile(true)
}

func (e *Environtment) GetConfig() map[string]interface{} {
	return e.configs
}

func (e *Environtment) loadFile(overload bool) {
	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range e.configs {
		if !currentEnv[key] || overload {
			val := fmt.Sprintf("%v", value)
			os.Setenv(key, val)
		}
	}
}

func Env(path string) Environtment {
	var environment Environtment
	environment = Environtment{Config(path)}
	environment.Load()

	return environment
}
