package utils

import "strconv"

const DEFAULT_PORT = 8000

func GetPort() int {
	port, err := strconv.Atoi(GetEnvWithDefault("PORT", strconv.Itoa(DEFAULT_PORT)))
	if err != nil {
		port = DEFAULT_PORT
	}

	return port
}
