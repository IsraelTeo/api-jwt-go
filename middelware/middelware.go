package middelware

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-jwt-go/auth"
)

func Log(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %q, Method %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

func SetMiddelwareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := auth.ValidateToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
