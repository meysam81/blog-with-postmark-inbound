package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/meysam81/x/config"
	"github.com/meysam81/x/sqlite"

	p "github.com/meysam81/tarzan/cmd/config"
)

type InboundEmail struct {
	From        string       `json:"From"`
	Subject     string       `json:"Subject"`
	HtmlBody    string       `json:"HtmlBody"`
	TextBody    string       `json:"TextBody"`
	Attachments []Attachment `json:"Attachments"`
}

type Attachment struct {
	Name        string `json:"Name"`
	Content     string `json:"Content"`
	ContentType string `json:"ContentType"`
	ContentID   string `json:"ContentID"`
}

type Post struct {
	ID          int
	Title       string
	Content     template.HTML
	AuthorEmail string
	CreatedAt   string
}

var (
	authorizedEmails = []string{"user1@example.com", "user3@example.com"}
	db               *sql.DB
	templates        *template.Template
	port             = "8000"
	rootPath         = "."
)

type appHandler struct {
	c *config.Config
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		rw := newResponseWriter(w)

		var reqHeaders []string
		for name, values := range r.Header {
			if strings.EqualFold(name, "Authorization") {
				continue
			}
			for _, value := range values {
				reqHeaders = append(reqHeaders, fmt.Sprintf("%s: %s", name, value))
			}
		}

		next(rw, r)

		var respHeaders []string
		for name, values := range rw.Header() {
			for _, value := range values {
				respHeaders = append(respHeaders, fmt.Sprintf("%s: %s", name, value))
			}
		}

		duration := time.Since(startTime)
		log.Printf(
			"%s %s %s %v %v - %d - %v",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			strings.Join(reqHeaders, ", "),
			strings.Join(respHeaders, ", "),
			rw.statusCode,
			duration,
		)
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	var err error

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.String(p.BASIC_AUTH_USERNAME) == "" {
		cfg.Set(p.BASIC_AUTH_USERNAME, "postmark")
	}
	if cfg.String(p.BASIC_AUTH_PASSWORD) == "" {
		cfg.Set(p.BASIC_AUTH_PASSWORD, "secret")
	}

	ctx := context.Background()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	if p := cfg.String(p.PORT); p != "" {
		port = p
	}

	dataPath, ok := os.LookupEnv("DATA_PATH")
	if ok && dataPath != "" {
		rootPath = dataPath
	}
	dbFilepath := filepath.Join(rootPath, "blog.db")
	staticDir := filepath.Join(rootPath, "static")

	db, err = sqlite.NewDB(ctx, dbFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        author_email TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("static/attachments", 0755)
	if err != nil {
		log.Fatalf("Error creating attachments directory: %v", err)
	}

	templates = template.Must(template.ParseFiles("templates/index.html"))

	app := &appHandler{
		c: cfg,
	}

	http.HandleFunc("/webhook", loggingMiddleware(app.webhookHandler))
	http.HandleFunc("/", loggingMiddleware(indexHandler))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	log.Println("Server started at :" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func (a *appHandler) webhookHandler(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok || username != a.c.String(p.BASIC_AUTH_USERNAME) || password != a.c.String(p.BASIC_AUTH_PASSWORD) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)

		return
	}

	var email InboundEmail
	if err := json.Unmarshal(body, &email); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	authorized := false
	for _, authEmail := range authorizedEmails {
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
			filepath := filepath.Join("static", "attachments", filename)

			if err := os.WriteFile(filepath, data, 0644); err != nil {
				log.Printf("Error saving attachment %s: %v", att.Name, err)
				continue
			}

			contentIDMap[att.ContentID] = "/static/attachments/" + filename
		}
	}

	for contentID, url := range contentIDMap {
		content = strings.ReplaceAll(content, "cid:"+contentID, url)
	}

	_, err = db.Exec("INSERT INTO posts (title, content, author_email) VALUES (?, ?, ?)",
		email.Subject, content, email.From)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, author_email, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Error querying posts", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorEmail, &p.CreatedAt); err != nil {
			log.Printf("Error scanning post: %v", err)
			continue
		}
		posts = append(posts, p)
	}

	if err := templates.ExecuteTemplate(w, "index.html", posts); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
