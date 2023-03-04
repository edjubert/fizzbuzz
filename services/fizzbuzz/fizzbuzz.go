package fizzbuzz

import (
	"io/ioutil"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/utils"
)

func SendFizzBuzz(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
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

	raw := FizzBuzz{}
	msg := raw.Generate(*params).String()

	utils.SendResponse(w, msg, constants.FIZZBUZZ, http.StatusOK, nil)
}
