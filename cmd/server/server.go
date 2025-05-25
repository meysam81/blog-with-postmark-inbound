package server

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/meysam81/x/httputils"

	"github.com/meysam81/tarzan/cmd/config"
	. "github.com/meysam81/tarzan/cmd/controllers"
	"github.com/meysam81/tarzan/cmd/datastore/sqlite"
)

func Main(frontend embed.FS) {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg.BaseUrl)

	authorizedEmails := []string{}

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	sqliteBuilder := &sqlite.Builder{}

	ds, err := sqliteBuilder.NewDatastore(ctxT, cfg)
	defer ds.Close()

	if err != nil {
		log.Fatalln("Error opening db:", err)
	}

	app := AppState{Config: cfg, AuthorizedEmails: authorizedEmails, DS: ds}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	err = os.MkdirAll(app.GetAttachmentPath(), 0755)
	if err != nil {
		log.Fatalf("Error creating attachments directory: %v", err)
	}

	if _, err := os.Stat(app.Config.AuthorizedEmailsPath); os.IsNotExist(err) {
		log.Fatalln("Missing authorized email filepath:", err)
	}
	content, err := os.ReadFile(app.Config.AuthorizedEmailsPath)
	if err != nil {
		log.Fatalln("Failed reading authorized emails", err)
	}
	err = json.Unmarshal(content, &app.AuthorizedEmails)
	if err != nil {
		log.Fatalln("Error reading authorized emails file:", err)
	}

	http.HandleFunc("/sitemap.xml", httputils.LoggingMiddleware(app.SitemapHandler))
	http.HandleFunc("/rss.xml", httputils.LoggingMiddleware(app.RSSHandler))

	http.HandleFunc("/api/posts", httputils.LoggingMiddleware(app.ListPosts))
	http.HandleFunc("/webhook", httputils.LoggingMiddleware(app.WebhookHandler))

	http.Handle("/api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir(app.Config.StoragePath))))
	distFS, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal("Error creating subdirectory filesystem:", err)
	}

	http.Handle("/", http.FileServer(http.FS(distFS)))

	log.Println("Server started at:", app.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), nil))
}
