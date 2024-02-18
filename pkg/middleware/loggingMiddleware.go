package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Infof(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
