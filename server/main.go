package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func apiIndex(w http.ResponseWriter, r *http.Request) {
	// handle GET /api request
}

func apiCreate(w http.ResponseWriter, r *http.Request) {
	// handle POST /api/create request
}

func apiUpdate(w http.ResponseWriter, r *http.Request) {
	// handle PUT /api/update request
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	// handle DELETE /api/delete request
}

func main() {
	r := chi.NewRouter()

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/", apiIndex)
		r.Post("/create", apiCreate)
		r.Put("/update", apiUpdate)
		r.Delete("/delete", apiDelete)
	})

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	r.FileServer("/", http.Dir(filesDir))

	http.ListenAndServe(":3000", r)
}
