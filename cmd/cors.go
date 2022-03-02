package cmd

import (
	"net/http"
)

func CustomCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("Access-Control-Allow-Origin", "*")

		rw.Header().Set("Access-Control-Allow-Headers", "Authorization")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			rw.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(rw, r)
	})
}
