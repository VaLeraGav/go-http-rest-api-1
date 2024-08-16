package middleware

import (
	"context"
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/errors"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/handlers"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func AuthenticateUser(store store.Store, sessionStore sessions.Store, sessionName string, ctxKeyUser int8) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := sessionStore.Get(r, sessionName)
			if err != nil {
				handlers.ErrorHandle(w, r, http.StatusInternalServerError, err)
				return
			}

			id, ok := session.Values["user_id"]
			if !ok {
				handlers.ErrorHandle(w, r, http.StatusUnauthorized, errors.NotAuthenticated)
				return
			}

			u, err := store.User().Find(id.(int))
			if err != nil {
				handlers.ErrorHandle(w, r, http.StatusUnauthorized, errors.NotAuthenticated)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
		})
	}
}
