package handlers

import (
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/model"
)

func Whoami(ctxKeyUser int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}
