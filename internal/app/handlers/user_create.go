package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/model"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
)

func HandelUserCreate(store store.Store) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			ErrorHandle(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := store.User().Create(u); err != nil {
			ErrorHandle(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		Respond(w, r, http.StatusCreated, u)
	}
}
