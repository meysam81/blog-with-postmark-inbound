package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *AppState) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := a.DS.ListPosts(r.Context(), maskEmail)
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

func (a *AppState) FetchPost(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postId")
	log.Println("Fetching post:", postIDStr)

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		log.Println("Invalid post ID:", err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := a.DS.FetchPost(r.Context(), postID, maskEmail)
	if err != nil {
		log.Println("Failed fetching post:", err)
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(post)
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
