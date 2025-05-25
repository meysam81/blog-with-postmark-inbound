package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/meysam81/tarzan/cmd/models"
)

func (a *AppState) WebhookHandler(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok || username != a.Config.AuthUsername || password != a.Config.AuthPassword {
		http.Error(w, "Unauthorized!", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Invalid request body provided", http.StatusBadRequest)
		return
	}

	var email models.InboundEmail
	if err := json.Unmarshal(body, &email); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Invalid json format in body", http.StatusBadRequest)
		return
	}

	authorized := false
	for _, authEmail := range a.AuthorizedEmails {
		if email.From == authEmail {
			authorized = true
			break
		}
	}
	if !authorized {
		log.Printf("Unauthorized sender: %s", email.From)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	content := email.HtmlBody
	if content == "" {
		content = email.TextBody
	}

	contentIDMap := make(map[string]string)
	for _, att := range email.Attachments {
		if att.ContentID != "" {

			data, err := base64.StdEncoding.DecodeString(att.Content)
			if err != nil {
				log.Printf("Error decoding attachment %s: %v", att.Name, err)
				continue
			}

			ext := ".bin"
			if strings.HasPrefix(att.ContentType, "image/png") {
				ext = ".png"
			} else if strings.HasPrefix(att.ContentType, "image/jpeg") {
				ext = ".jpg"
			}

			filename := uuid.New().String() + ext
			filepath := filepath.Join(a.GetAttachmentPath(), filename)

			if err := os.WriteFile(filepath, data, 0644); err != nil {
				log.Printf("Error saving attachment %s: %v", att.Name, err)
				http.Error(w, "Error occurred while saving your attachments", http.StatusBadRequest)
				continue
			}

			contentIDMap[att.ContentID] = "/api/assets/attachments/" + filename
		}
	}

	for contentID, url := range contentIDMap {
		content = strings.ReplaceAll(content, "cid:"+contentID, url)
	}

	err = a.DS.Insert(r.Context(), &email)
	if err != nil {
		log.Println("Failed inserting to datastore:", err)
		http.Error(w, "Failed creating post:", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
