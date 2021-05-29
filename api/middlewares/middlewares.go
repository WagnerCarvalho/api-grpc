package middlewares

import (
	"api-grpc/api/restutils"
	"log"
	"net/http"
	"strings"
	"time"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(`{"proto": "%s", "method": "%s", "route": "%s%s", request_time: "%v"}`, r.Proto, r.Method, r.Host, r.URL.Path, time.Since(t))
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header := strings.TrimSpace(r.Header.Get("Authorization"))
		splitred := strings.Split(header, " ")
		if len(splitred) != 2 {
			restutils.WriteError(w, http.StatusUnauthorized, restutils.ErrUnauthorized)
		}
	}
}
