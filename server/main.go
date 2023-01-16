package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func apiIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"text": "I am a GET man."})
}

func apiCreate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"text": "I am a POST man."})
}

func apiUpdate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"text": "I am a PUT man."})
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"text": "I am a DELETE man."})
}

func main() {
	r := chi.NewRouter()

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Use(setContentType)
		r.Get("/", apiIndex)
		r.Post("/", apiCreate)
		r.Put("/", apiUpdate)
		r.Delete("/", apiDelete)
	})

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	r.Handle("/*", http.StripPrefix("/", http.FileServer(filesDir)))

	http.ListenAndServe(":80", r)
}

func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
