package assets

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Mount(r chi.Router) {
	fs := http.Dir("assets/dist")
	fileServer := http.FileServer(fs)

	r.Route("/dist", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "no-store")
				next.ServeHTTP(w, r)
			})
		})
		r.Handle("/*", http.StripPrefix("/dist", fileServer))
	})
}
