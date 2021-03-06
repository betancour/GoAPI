package app

import (
	"net/http"
	u "../utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resources was not found in our server"))
		next.ServeHTTP(w, r)
	})
}
