package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/meysam81/tarzan/cmd/metrics"
	"github.com/meysam81/tarzan/cmd/models"
	"github.com/russross/blackfriday/v2"
)

func (a *AppState) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r.Response != nil {
			metrics.RecordEmailProcessed(r.Response.Status)
		}
	}()

	now := time.Now()
	defer func() { metrics.RecordEmailProcessingDuration(float64(time.Since(now))) }()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Invalid request body provided", http.StatusBadRequest)
		return
	}

	payloadName := strings.ReplaceAll(time.Now().Format(time.DateTime), " ", "T")
	a.Filestore.Save(payloadName, body)

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
	if email.TextBody != "" && isMarkdown(email.TextBody) {
		content = convertMarkdownToHtml(email.TextBody)
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

			data, err := base64.StdEncoding.DecodeString(att.Content)
			if err != nil {
				log.Println("Failed decoding base64 content:", err)
				continue
			}

			filename := uuid.New().String() + ext

			err = a.Filestore.Save(filename, data)
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
	select {
	case *a.Signal <- 1:
		log.Println("Signal sent successfully")
	case <-time.After(100 * time.Millisecond):
		log.Println("Signal send timed out - no receiver available")
	}

	metrics.IncrementPostsTotal()

	w.WriteHeader(http.StatusNoContent)
}

func convertMarkdownToHtml(content string) string {
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.HrefTargetBlank,
	})

	extensions := blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
	htmlContent := blackfriday.Run([]byte(content), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extensions))
	content = string(htmlContent)

	return content
}

func isMarkdown(content string) bool {
	markdownIndicators := []string{"#", "*", "_", "`", "[", "](", "```"}
	for _, indicator := range markdownIndicators {
		if strings.Contains(content, indicator) {
			return true
		}
	}
	return false
}
