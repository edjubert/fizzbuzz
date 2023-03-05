package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/services/fizzbuzz"
	"github.com/edjubert/leboncoin/utils"
)

func FizzBuzzPost(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SendResponse(w, "", constants.FIZZBUZZ, http.StatusInternalServerError, err)
		return
	}

	params, err := utils.UnmarshalParams(body)
	if err != nil {
		utils.SendResponse(w, "", constants.FIZZBUZZ, http.StatusInternalServerError, err)
		return
	}

	msg := fizzbuzz.GetFizzBuzz(*params)
	utils.SendResponse(w, msg, constants.FIZZBUZZ, http.StatusOK, nil)
}

func FizzBuzz(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	switch r.Method {
	case "POST":
		FizzBuzzPost(w, r, redis)
	}
}
