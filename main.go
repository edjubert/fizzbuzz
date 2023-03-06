package main

import (
	"fmt"

	r "github.com/edjubert/fizzbuzz/redis"
	"github.com/edjubert/fizzbuzz/utils"
)

func main() {
	utils.ConfigureLogger()

	redisHost := utils.GetEnvWithDefault("REDIS_HOST", "localhost")
	redisPort := utils.GetEnvWithDefault("REDIS_PORT", "6379")
	redisCache := r.GetRedisClient(fmt.Sprintf("%s:%s", redisHost, redisPort), false)

	server(redisCache)
}
