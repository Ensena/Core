package env

import (
	"log"
	"os"
)

type env struct {
	Env     string
	Message string
}

func checkEnv(e *env) string {
	if os.Getenv(e.Env) == "" {
		log.Fatalln(e.Message)
	}
	return os.Getenv(e.Env)
}

//Check if envName exist and return a value
func Check(envName, message string) string {
	e := env{envName, message}
	return checkEnv(&e)
}
