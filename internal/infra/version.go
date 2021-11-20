package infra

import (
	"net/http"
)

var Version string // do not remove or modify

func VersionHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-api-version", Version)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}
