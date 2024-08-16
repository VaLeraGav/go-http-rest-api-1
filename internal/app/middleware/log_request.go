package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func LogRequest(logger *logrus.Logger, ctxKeyRequestID int8) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := logger.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"request_id":  r.Context().Value(ctxKeyRequestID),
			})
			logger.Infof("started %s %s", r.Method, r.RequestURI)

			start := time.Now()

			// преопределили
			rw := &responseWriter{w, http.StatusOK}

			// http.ResponseWriter - передаем интерфес и он не имеет доступа к коду ответа
			next.ServeHTTP(rw, r)

			var level logrus.Level
			switch {
			case rw.code >= 500:
				level = logrus.ErrorLevel
			case rw.code >= 400:
				level = logrus.WarnLevel
			default:
				level = logrus.InfoLevel
			}
			logger.Logf(level, "completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
		})
	}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
