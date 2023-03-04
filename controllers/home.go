package controllers

import (
	"net/http"

	"github.com/edjubert/leboncoin/services/home"
)

func Base(w http.ResponseWriter, r *http.Request) {
	home.Base(w, r)
}
