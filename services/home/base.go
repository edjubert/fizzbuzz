package home

import (
	"net/http"

	"github.com/edjubert/leboncoin/utils"
)

func Base(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "ok", "/", http.StatusOK, nil)
}
