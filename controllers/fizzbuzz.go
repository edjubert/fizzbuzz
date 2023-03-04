package controllers

import (
	"net/http"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/services/fizzbuzz"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	switch r.Method {
	case "POST":
		fizzbuzz.SendFizzBuzz(w, r, redis)
	}
}
