package config

import (
	"os"
	"strings"
)

func GetEnv(name string) interface{} {
	switch name {
	case "addr":
		addr := os.Getenv(strings.ToUpper(name))
		if addr == "" {
			addr = ":3000"
		}
		return addr
	default:
		return os.Getenv(strings.ToUpper(name))
	}
}
