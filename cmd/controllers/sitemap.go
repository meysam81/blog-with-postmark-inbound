package controllers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

func (a *AppState) SitemapHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := a.DS.ListPosts(r.Context())
	if err != nil {
		http.Error(w, "Error generating sitemap", http.StatusInternalServerError)
		return
	}

	baseURL := a.Config.BaseUrl
	if host := r.Header.Get("Host"); host != "" {
		scheme := "https"
		if r.TLS == nil {
			scheme = "http"
		}
		baseURL = fmt.Sprintf("%s://%s", scheme, host)
	}

	urlset := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs: []URL{
			{
				Loc:        baseURL,
				LastMod:    time.Now().Format(time.DateOnly),
				ChangeFreq: "daily",
				Priority:   "1.0",
			},
		},
	}

	for _, post := range *posts {
		lastMod := ""
		if parsedTime, err := time.Parse(time.DateTime, post.CreatedAt); err == nil {
			lastMod = parsedTime.Format(time.DateOnly)
		} else if parsedTime, err := time.Parse(time.RFC3339, post.CreatedAt); err == nil {
			lastMod = parsedTime.Format(time.DateOnly)
		}

		postURL := URL{
			Loc:        fmt.Sprintf("%s/post/%d", baseURL, post.ID),
			LastMod:    lastMod,
			ChangeFreq: "monthly",
			Priority:   "0.8",
		}
		urlset.URLs = append(urlset.URLs, postURL)
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)

	xmlData, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		http.Error(w, "Error generating XML", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return
	}
	if _, err := w.Write(xmlData); err != nil {
		return
	}
}
