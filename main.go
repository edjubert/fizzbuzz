package main

import (
	"fmt"

	r "github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/utils"
)

func main() {
	utils.ConfigureLogger()

	redisHost := utils.GetEnvWithDefault("REDIS_HOST", "localhost")
	redisPort := utils.GetEnvWithDefault("REDIS_PORT", "6379")
	redisCache := r.GetRedisClient(fmt.Sprintf("%s:%s", redisHost, redisPort), false)

	server(redisCache)
}
