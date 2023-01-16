package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func apiIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello I am a GET man."))
}

func apiCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello I am a POST man."))
}

func apiUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello I am a PUT man."))
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello I am a DELETE man."))
}

func main() {
	r := chi.NewRouter()

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/", apiIndex)
		r.Get("/get", apiIndex)
		r.Post("/create", apiCreate)
		r.Put("/update", apiUpdate)
		r.Delete("/delete", apiDelete)
	})

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	r.Handle("/*", http.StripPrefix("/", http.FileServer(filesDir)))

	http.ListenAndServe(":3000", r)
}
