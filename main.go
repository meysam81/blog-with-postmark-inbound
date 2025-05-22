package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	p "github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/x/config"
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
	redisClient      *redis.Client
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
		err = cfg.Set(p.BASIC_AUTH_USERNAME, "postmark")
		if err != nil {
			log.Fatalln(err)
		}
	}
	if cfg.String(p.BASIC_AUTH_PASSWORD) == "" {
		err = cfg.Set(p.BASIC_AUTH_PASSWORD, "secret")
		if err != nil {
			log.Fatalln(err)
		}
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
	staticDir := filepath.Join(rootPath, "static")

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:       redisAddr,
		ClientName: fmt.Sprintf("tarzan-%s", runtime.Version()),
		Password:   "", // No password set
		DB:         0,  // Default DB
	})
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}()

	// Check Redis connection
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
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

	ctx := context.Background()
	postID, err := redisClient.Incr(ctx, "post:counter").Result()
	if err != nil {
		log.Printf("Error incrementing post counter: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	postKey := fmt.Sprintf("post:%d", postID)
	postData := map[string]interface{}{
		"title":        email.Subject,
		"content":      content,
		"author_email": email.From,
		"created_at":   time.Now().Format(time.RFC3339),
	}

	if err := redisClient.HSet(ctx, postKey, postData).Err(); err != nil {
		log.Printf("Error inserting post: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Add post ID to a sorted set for ordering by creation time
	if err := redisClient.ZAdd(ctx, "posts", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Err(); err != nil {
		log.Printf("Error adding post to sorted set: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	postIDs, err := redisClient.ZRevRange(ctx, "posts", 0, -1).Result()
	if err != nil {
		http.Error(w, "Error querying posts", http.StatusInternalServerError)
		return
	}

	var posts []Post
	for i, idStr := range postIDs {
		postKey := fmt.Sprintf("post:%s", idStr)
		postData, err := redisClient.HGetAll(ctx, postKey).Result()
		if err != nil {
			log.Printf("Error fetching post %s: %v", idStr, err)
			continue
		}

		post := Post{
			ID:          i + 1,
			Title:       postData["title"],
			Content:     template.HTML(postData["content"]),
			AuthorEmail: postData["author_email"],
			CreatedAt:   postData["created_at"],
		}
		posts = append(posts, post)
	}

	if err := templates.ExecuteTemplate(w, "index.html", posts); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
