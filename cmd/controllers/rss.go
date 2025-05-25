package routes

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	AtomNS  string   `xml:"xmlns:atom,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Language      string   `xml:"language,omitempty"`
	PubDate       string   `xml:"pubDate,omitempty"`
	LastBuildDate string   `xml:"lastBuildDate,omitempty"`
	Generator     string   `xml:"generator,omitempty"`
	AtomLink      AtomLink `xml:"atom:link"`
	Items         []Item   `xml:"item"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate,omitempty"`
	GUID        string `xml:"guid"`
}

func (a *AppState) RSSHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := a.DS.List(r.Context())
	if err != nil {
		http.Error(w, "Error generating RSS feed", http.StatusInternalServerError)
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

	channel := Channel{
		Title:         "Tarzan Blog",
		Link:          baseURL,
		Description:   "Latest posts from Tarzan Blog",
		Language:      "en-us",
		Generator:     "Tarzan RSS Generator",
		LastBuildDate: time.Now().Format(time.RFC1123Z),
		AtomLink: AtomLink{
			Href: fmt.Sprintf("%s/rss", baseURL),
			Rel:  "self",
			Type: "application/rss+xml",
		},
		Items: []Item{},
	}

	for _, post := range *posts {
		pubDate := ""
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt); err == nil {
			pubDate = parsedTime.Format(time.RFC1123Z)
		} else if parsedTime, err := time.Parse(time.RFC3339, post.CreatedAt); err == nil {
			pubDate = parsedTime.Format(time.RFC1123Z)
		}

		item := Item{
			Title:       post.Title,
			Link:        fmt.Sprintf("%s/post/%d", baseURL, post.ID),
			Description: post.Content,
			PubDate:     pubDate,
			GUID:        fmt.Sprintf("%s/post/%d", baseURL, post.ID),
		}
		channel.Items = append(channel.Items, item)
	}

	rss := RSS{
		Version: "2.0",
		AtomNS:  "http://www.w3.org/2005/Atom",
		Channel: channel,
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	xmlData, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		http.Error(w, "Error generating RSS XML", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return
	}
	if _, err := w.Write(xmlData); err != nil {
		return
	}
}
