package http

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

var versionedAssetMatcher = regexp.MustCompile(`([^.]+)\.[a-z0-9]+(\.(?:js|css))`)

func VersionedAssets(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if versionedAssetMatcher.MatchString(r.URL.Path) {
			r.URL.Path = versionedAssetMatcher.ReplaceAllString(r.URL.Path, `$1$2`)
		}

		next.ServeHTTP(w, r)
	})
}

func Static(mux chi.Router) {
	staticHandler := http.FileServer(http.Dir("public"))
	mux.Get(`/{:[^.]+\.[^.]+}`, staticHandler.ServeHTTP)
	mux.Get(`/{:images|scripts|styles}/*`, staticHandler.ServeHTTP)
}
