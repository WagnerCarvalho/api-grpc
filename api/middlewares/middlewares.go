package middlewares

import (
	"api-grpc/api/restutils"
	"api-grpc/security"
	"log"
	"net/http"
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
		tokenString, err := security.ExtractToken(r)
		if err != nil {
			restutils.WriteError(w, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return
		}

		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("Error on parsed token:", err.Error())
			restutils.WriteError(w, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("Token invalid", tokenString)
			restutils.WriteError(w, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return
		}

		next(w, r)

	}
}
