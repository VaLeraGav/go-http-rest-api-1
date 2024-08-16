package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/errors"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
	"github.com/gorilla/sessions"
)

func SessionsCreate(store store.Store, sessionStore sessions.Store, sessionName string) http.HandlerFunc {
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

		u, err := store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			ErrorHandle(w, r, http.StatusUnauthorized, errors.IncorrectEmailOrPassword)
			return
		}

		session, err := sessionStore.Get(r, sessionName)
		if err != nil {
			ErrorHandle(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := sessionStore.Save(r, w, session); err != nil {
			ErrorHandle(w, r, http.StatusInternalServerError, err)
			return
		}

		Respond(w, r, http.StatusOK, nil)
	}
}
