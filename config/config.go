package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// New returns map[string]string of .env vars
func New(filename string) map[string]string {
	_, b, _, _ := runtime.Caller(1)
	envFilePath := path.Join(path.Dir(b), filename)

	envMap, err := godotenv.Read(envFilePath)
	if err != nil {
		panic(fmt.Errorf("fatal error: %s", err))
	}

	return envMap
}
