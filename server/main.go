package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func spaHandler(distDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(distDir, r.URL.Path)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			http.ServeFile(w, r, path)
			return
		}
		http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
	}
}

func main() {
	http.HandleFunc("/", spaHandler("web/dist"))
	http.ListenAndServe(":2222", nil)
}
