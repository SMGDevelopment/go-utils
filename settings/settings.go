package settings

import (
	"os"
)

func EnvVar(env string, fallback string) string {
	v, ok := os.LookupEnv(env)
	if ok {
		return v
	}
	return fallback
}
