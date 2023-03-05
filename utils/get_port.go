package utils

import "strconv"

const DEFAULT_PORT = 8000
const DEFAULT_HEALTH_CHECK_PORT = 8001

func GetPort(healthCheck bool) int {
	port := DEFAULT_PORT
	var err error
	if healthCheck {
		port, err = strconv.Atoi(GetEnvWithDefault("HEALT_CHECK_PORT", strconv.Itoa(DEFAULT_HEALTH_CHECK_PORT)))
	} else {
		port, err = strconv.Atoi(GetEnvWithDefault("PORT", strconv.Itoa(DEFAULT_PORT)))
	}
	if err != nil {
		port = DEFAULT_PORT
	}

	return port
}
