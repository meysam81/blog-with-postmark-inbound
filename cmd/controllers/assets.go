package controllers

import (
	"net/http"
	"path/filepath"
	"strings"
)

func (a *AppState) AssetsHandler(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/api/assets/attachments/")
	if filename == "" {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	data, err := a.Filestore.LoadBytes(filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	ext := filepath.Ext(filename)
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".pdf":
		w.Header().Set("Content-Type", "application/pdf")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
