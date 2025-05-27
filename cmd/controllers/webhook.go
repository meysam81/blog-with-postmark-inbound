package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/meysam81/tarzan/cmd/metrics"
	"github.com/meysam81/tarzan/cmd/models"
)

func (a *AppState) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { metrics.RecordEmailProcessed(r.Response.Status) }()

	now := time.Now()
	defer func() { metrics.RecordEmailProcessingDuration(float64(time.Since(now))) }()

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
	content = email.HtmlBody
	if content == "" {
		content = email.TextBody
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

	finalContent := processContentForMarkdown(email.TextBody, email.HtmlBody, contentIDMap)

	err = a.DS.InsertEmail(r.Context(), &models.EmailInsertDB{
		From:     email.From,
		FromName: email.FromName,
		Subject:  email.Subject,
		Content:  finalContent,
	})
	if err != nil {
		log.Println("Failed inserting to datastore:", err)
		http.Error(w, "Failed creating post:", http.StatusBadRequest)
		return
	}

	log.Println("Notifying the signal for updated list...")
	*a.Signal <- 1

	metrics.IncrementPostsTotal()

	w.WriteHeader(http.StatusOK)
}

func processContentForMarkdown(textBody, htmlBody string, contentIDMap map[string]string) string {
	if textBody == "" {
		return replaceContentIDs(htmlBody, contentIDMap)
	}

	if htmlBody == "" {
		return textBody
	}

	imageMarkdown := extractImagesAsMarkdown(htmlBody, contentIDMap)
	if imageMarkdown == "" {
		return textBody
	}

	enhancedText := insertImagesIntoText(textBody, imageMarkdown)
	return enhancedText
}

func extractImagesAsMarkdown(htmlBody string, contentIDMap map[string]string) string {
	imgRegex := regexp.MustCompile(`<img[^>]+src="cid:([^"]+)"[^>]*alt="([^"]*)"[^>]*>`)
	matches := imgRegex.FindAllStringSubmatch(htmlBody, -1)

	var markdownImages []string
	for _, match := range matches {
		if len(match) >= 3 {
			contentID := match[1]
			altText := match[2]

			if url, exists := contentIDMap[contentID]; exists {
				markdownImg := "![" + altText + "](" + url + ")"
				markdownImages = append(markdownImages, markdownImg)
			}
		}
	}

	return strings.Join(markdownImages, "\n\n")
}

func insertImagesIntoText(textBody, imageMarkdown string) string {
	lines := strings.Split(textBody, "\n")
	var result []string
	imageInserted := false

	for i, line := range lines {
		result = append(result, line)

		if !imageInserted && (strings.HasPrefix(line, "---") ||
			strings.Contains(strings.ToLower(line), "photo by") ||
			strings.Contains(strings.ToLower(line), "image") ||
			(i > 0 && i < 5 && strings.TrimSpace(line) == "")) {

			result = append(result, "", imageMarkdown, "")
			imageInserted = true
		}
	}

	if !imageInserted && imageMarkdown != "" {
		titleEndIndex := -1
		for i, line := range lines {
			if strings.HasPrefix(line, "#") {
				titleEndIndex = i
				break
			}
		}

		if titleEndIndex >= 0 && titleEndIndex < len(result)-1 {
			insertIndex := titleEndIndex + 1
			newResult := make([]string, 0, len(result)+3)
			newResult = append(newResult, result[:insertIndex]...)
			newResult = append(newResult, "", imageMarkdown, "")
			newResult = append(newResult, result[insertIndex:]...)
			result = newResult
		}
	}

	return strings.Join(result, "\n")
}

func replaceContentIDs(content string, contentIDMap map[string]string) string {
	for contentID, url := range contentIDMap {
		content = strings.ReplaceAll(content, "cid:"+contentID, url)
	}
	return content
}
