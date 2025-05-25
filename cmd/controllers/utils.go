package controllers

import (
	"strings"

	"github.com/meysam81/tarzan/cmd/models"
	"github.com/russross/blackfriday/v2"
)

func maskEmail(p *models.Post) {
	at := strings.Index(p.AuthorEmail, "@")
	if at <= 1 {
		p.AuthorEmail = "***"
		return
	}
	name := p.AuthorEmail[:at]
	domain := p.AuthorEmail[at:]
	if len(name) <= 2 {
		p.AuthorEmail = string(name[0]) + "*" + domain
		return
	}
	p.AuthorEmail = string(name[0]) + strings.Repeat("*", len(name)-2) + string(name[len(name)-1]) + domain
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
