package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/meysam81/x/httputils"
	"github.com/meysam81/x/sqlite"

	"github.com/meysam81/tarzan/cmd/config"
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
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorEmail string `json:"author-email"`
	CreatedAt   string `json:"created-at"`
}

type appHandler struct {
	cfg              *config.Config
	authorizedEmails []string
	db               *sql.DB
}

func (a *appHandler) getAttachmentPath() string {
	return filepath.Join(a.cfg.StoragePath, "attachments")
}

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	authorizedEmails := []string{}

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	db, err := sqlite.NewDB(ctxT, cfg.DbPath, sqlite.WithJournalMode(""), sqlite.WithMode(""))
	if err != nil {
		log.Fatalln("Error opening db:", err)
	}

	app := appHandler{cfg, authorizedEmails, db}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

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

	err = os.MkdirAll(app.getAttachmentPath(), 0755)
	if err != nil {
		log.Fatalf("Error creating attachments directory: %v", err)
	}

	if _, err := os.Stat(app.cfg.AuthorizedEmailsPath); os.IsNotExist(err) {
		log.Fatalln("Missing authorized email filepath:", err)
	}
	content, err := os.ReadFile(app.cfg.AuthorizedEmailsPath)
	if err != nil {
		log.Fatalln("Failed reading authorized emails", err)
	}
	err = json.Unmarshal(content, &app.authorizedEmails)
	if err != nil {
		log.Fatalln("Error reading authorized emails file:", err)
	}

	http.HandleFunc("/api/posts", httputils.LoggingMiddleware(app.listPosts))
	http.HandleFunc("/webhook", httputils.LoggingMiddleware(app.webhookHandler))
	http.Handle("/api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir(app.cfg.StoragePath))))
	distFS, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal("Error creating subdirectory filesystem:", err)
	}
	http.Handle("/", http.FileServer(http.FS(distFS)))

	log.Println("Server started at:", app.cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.cfg.Port), nil))
}

func (a *appHandler) webhookHandler(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok || username != a.cfg.AuthUsername || password != a.cfg.AuthPassword {
		http.Error(w, "Unauthorized!", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Invalid request body provided", http.StatusBadRequest)
		return
	}

	var email InboundEmail
	if err := json.Unmarshal(body, &email); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Invalid json format in body", http.StatusBadRequest)
		return
	}

	authorized := false
	for _, authEmail := range a.authorizedEmails {
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
			filepath := filepath.Join(a.getAttachmentPath(), filename)

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

	_, err = a.db.Exec("INSERT INTO posts (title, content, author_email) VALUES (?, ?, ?)",
		email.Subject, content, email.From)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		http.Error(w, "Error saving your post to the database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func maskEmail(p *Post) {
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

func (a *appHandler) listPosts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()
	rows, err := a.db.QueryContext(ctxT, "SELECT id, title, content, author_email, created_at FROM posts ORDER BY created_at DESC")
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
		maskEmail(&p)
		posts = append(posts, p)
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
