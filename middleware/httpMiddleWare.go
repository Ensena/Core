package middleware

import (
	"net/http"
)

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("AuthorizationToken")
		if Validate(token) == false {
			w.WriteHeader(403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
