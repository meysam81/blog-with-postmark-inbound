package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *AppState) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := a.DS.List(r.Context(), maskEmail)
	if err != nil {
		log.Println("Failed listing posts:", err)
		http.Error(w, "Failed retrieving posts", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&posts)
	if err != nil {
		log.Println("Failed marshaling to json:", err)
		http.Error(w, "Failed responding to request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Println("Failed writing to response writer:", err)
	}
}
