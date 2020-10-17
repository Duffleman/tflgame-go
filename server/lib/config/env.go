package config

import (
	"os"
	"strings"
)

func GetEnv(name string) interface{} {
	return os.Getenv(strings.ToUpper(name))
}
