//go:build production

package assets

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// The production build embeds the generated stylesheet into the binary.
//
//go:embed dist/tailwind.css
var embeddedAssets embed.FS

func Mount(r chi.Router) {
	dist, err := fs.Sub(embeddedAssets, "dist")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(dist))

	r.Route("/dist", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
				next.ServeHTTP(w, r)
			})
		})
		r.Handle("/*", http.StripPrefix("/dist", fileServer))
	})
}
