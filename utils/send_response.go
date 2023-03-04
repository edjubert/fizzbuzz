package utils

import (
	"net/http"

	"github.com/gookit/slog"
)

func SendResponse(w http.ResponseWriter, msg, addr string, status int, err error) {
	if err != nil {
		slog.Errorf("[%s] -> %d: %s", addr, status, err.Error())
	} else {
		slog.Infof("[%s] -> %d", addr, status)
		slog.Tracef("%s", msg)
	}

	w.WriteHeader(status)
	w.Write([]byte(msg))
}
