package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read basic auth information
		usr, _, ok := r.BasicAuth()

		// if there is no basic auth (no matter which credentials)
		if !ok {
			errMsg := "Authentication error!"
			// return a 403 forbidden
			http.Error(w, errMsg, http.StatusForbidden)
			log.Println(errMsg)

			// stop processing route
			return
		}

		log.Printf("User %s logged in.", usr)
		// if everything is ok, call next middleware or handler
		next.ServeHTTP(w, r)
	})
}
