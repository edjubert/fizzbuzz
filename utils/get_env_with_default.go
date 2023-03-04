package utils

import "os"

func GetEnvWithDefault(env, defaultValue string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}
