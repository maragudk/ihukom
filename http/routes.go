package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maragudk/httph"
)

func (s *Server) setupRoutes() {
	s.mux.Group(func(r chi.Router) {
		r.Use(middleware.Compress(5))
		r.Use(middleware.RealIP)

		r.Group(func(r chi.Router) {
			r.Use(VersionedAssets)

			Static(r)
		})

		r.Group(func(r chi.Router) {
			r.Use(httph.NoClickjacking, httph.ContentSecurityPolicy(func(opts *httph.ContentSecurityPolicyOptions) {
				opts.ConnectSrc = "'self'"
				opts.ImgSrc = "'self'"
				opts.ManifestSrc = "'self'"
				opts.ScriptSrc = "'self'"
			}))
			r.Use(middleware.SetHeader("Content-Type", "text/html; charset=utf-8"))

			Home(r, s.db)
			Notes(r, s.db)
		})
	})
}
