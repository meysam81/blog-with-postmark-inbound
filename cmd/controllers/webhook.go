package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

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
	if a.Config.DangerouslyAcceptAllSenders {
		authorized = true
	} else {
		verifiedEmail, err := a.DS.FindAuthorizedSenderByEmail(r.Context(), email.From)
		if err != nil {
			log.Println("The sender is unauthorized", email.From, err)
		}

		if verifiedEmail != nil {
			authorized = true
		}
	}

	if !authorized {
		log.Printf("Unauthorized sender: %s", email.From)
		http.Error(w, "You are not allowed to publish posts", http.StatusForbidden)
		return
	}

	var content string
	content = email.TextBody
	if content != "" {
		if isMarkdown(content) {
			content = convertMarkdownToHtml(content)
		}
	} else {
		content = email.HtmlBody
	}

	contentIDMap := make(map[string]string)
	for _, att := range email.Attachments {
		if att.ContentID != "" {

			ext := ".bin"
			if strings.HasPrefix(att.ContentType, "image/png") {
				ext = ".png"
			} else if strings.HasPrefix(att.ContentType, "image/jpeg") {
				ext = ".jpg"
			}

			filename, err := a.Filestore.Save(att.Content, ext)
			if err != nil {
				log.Println("Failed saving attachment:", err)
				continue
			}

			contentIDMap[att.ContentID] = "/api/assets/attachments/" + filename
		}
	}

	for contentID, url := range contentIDMap {
		content = strings.ReplaceAll(content, "cid:"+contentID, url)
	}

	err = a.DS.InsertEmail(r.Context(), &models.EmailInsertDB{
		From:     email.From,
		FromName: email.FromName,
		Subject:  email.Subject,
		Content:  content,
	})
	if err != nil {
		log.Println("Failed inserting to datastore:", err)
		http.Error(w, "Failed creating post:", http.StatusBadRequest)
		return
	}

	log.Println("Notifying the signal for updated list...")
	*a.Signal <- 1

	w.WriteHeader(http.StatusOK)
}
