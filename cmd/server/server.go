package server

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/meysam81/x/httputils"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/controllers"
	"github.com/meysam81/tarzan/cmd/datastore/sqlite"
	"github.com/meysam81/tarzan/cmd/filestore/filesystem"
)

func Main(frontend embed.FS) {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	sqliteBuilder := &sqlite.Builder{}
	filesystemBuilder := &filesystem.FilesystemBuilder{}

	ds, err := sqliteBuilder.NewDatastore(ctxT, cfg)
	defer func() {
		if err := ds.Close(); err != nil {
			log.Printf("Error closing datastore: %v", err)
		}
	}()

	if err != nil {
		log.Fatalln("Error opening db:", err)
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins for simplicity
	}
	signal := make(chan uint8)

	ctxT, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	filestore, err := filesystemBuilder.NewFilestore(ctxT, cfg)
	if err != nil {
		log.Fatalln("Failed initializing filestore")
	}

	app := controllers.AppState{Config: cfg, DS: ds, Upgrader: upgrader, Signal: &signal, Filestore: filestore}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/sitemap.xml", httputils.LoggingMiddleware(app.SitemapHandler))
	r.Get("/rss.xml", httputils.LoggingMiddleware(app.RSSHandler))

	r.Route("/api", func(r chi.Router) {
		r.Get("/posts", httputils.LoggingMiddleware(app.ListPosts))
		r.Handle("/assets/*", http.StripPrefix("/api/assets/", http.FileServer(http.Dir(app.Config.StoragePath))))
	})

	r.Post("/webhook", httputils.LoggingMiddleware(app.WebhookHandler))

	r.Get("/ws", app.Websocket)

	distFS, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal("Error creating subdirectory filesystem:", err)
	}

	// SPA hack: serve the static files first and fallback to index.html
	fileServer := http.FileServer(http.FS(distFS))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		file, err := distFS.Open(r.URL.Path[1:])
		if err == nil {
			defer file.Close()
			if stat, err := file.Stat(); err == nil && !stat.IsDir() {
				http.ServeContent(w, r, r.URL.Path, stat.ModTime(), file.(io.ReadSeeker))
				return
			}
		}

		indexFile, err := distFS.Open("index.html")
		defer indexFile.Close()

		if err != nil {
			log.Println("Error reading index.html", err)
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}

		stat, err := indexFile.Stat()
		if err != nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeContent(w, r, "/", stat.ModTime(), indexFile.(io.ReadSeeker))
	})

	log.Println("Server started at:", app.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), r))
}
