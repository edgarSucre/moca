package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func loggerMiddleware(logger *log.Entry) middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			log.Infof("%s=%s, %s=%s, %s=%s",
				"Host", r.Host,
				"URL", r.URL,
				"Method", r.Method,
			)

			next(w, r)
		}
	}
}
