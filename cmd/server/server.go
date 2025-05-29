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
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/controllers"
	"github.com/meysam81/tarzan/cmd/datastore"
	"github.com/meysam81/tarzan/cmd/datastore/redis"
	"github.com/meysam81/tarzan/cmd/datastore/sqlite"
	"github.com/meysam81/tarzan/cmd/filestore"
	"github.com/meysam81/tarzan/cmd/filestore/filesystem"
	redisFilestore "github.com/meysam81/tarzan/cmd/filestore/redis"
	mw "github.com/meysam81/tarzan/cmd/middleware"
)

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net; style-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net https://fonts.googleapis.com; font-src https://fonts.gstatic.com data:; img-src 'self' data: blob:; connect-src 'self' wss: ws:; manifest-src 'self'; frame-ancestors 'none'; report-uri /csp-violation-report")
		next.ServeHTTP(w, r)
	})
}

func Main(frontend embed.FS) {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	datastoreBuilder := map[string]datastore.BuildDatastore{
		"sqlite": &sqlite.Builder{},
		"redis":  &redis.Builder{},
	}[cfg.DataStore]

	filesystemBuilder := map[string]filestore.BuildFilestore{
		"filesystem": &filesystem.Builder{},
		"redis":      &redisFilestore.Builder{},
	}[cfg.FileStore]

	ds, err := datastoreBuilder.NewDatastore(ctxT, cfg)
	if err != nil {
		log.Fatalln("Error opening db:", err)
	}
	defer func() {
		if ds != nil {
			if err := ds.Close(); err != nil {
				log.Printf("Error closing datastore: %v", err)
			}
		}
	}()

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

	r.Use(securityHeaders)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	promRouter := chi.NewRouter()
	promRouter.Use(mw.PrometheusMiddleware)

	promRouter.Get("/sitemap.xml", httputils.LoggingMiddleware(app.SitemapHandler))
	promRouter.Get(app.Config.RSSEndpoint, httputils.LoggingMiddleware(app.RSSHandler))

	if app.Config.MetricsEnabled {
		if app.Config.MetricsAuthEnabled {
			metricsCreds := map[string]string{
				app.Config.MetricsAuthUsername: app.Config.MetricsAuthPassword,
			}
			promRouter.With(middleware.BasicAuth("Metrics Authentication", metricsCreds)).Get("/metrics", promhttp.Handler().ServeHTTP)
		} else {
			promRouter.Get("/metrics", promhttp.Handler().ServeHTTP)
		}
	}

	promRouter.Route("/api", func(r chi.Router) {
		r.Get("/posts", httputils.LoggingMiddleware(app.ListPosts))
		r.Get("/posts/{postId}", httputils.LoggingMiddleware(app.FetchPost))
		r.Handle("/assets/*", httputils.LoggingMiddleware(app.AssetsHandler))
	})

	promRouter.Post("/csp-violation-report", app.CSPViolationHandler)

	webhookCreds := map[string]string{
		app.Config.AuthUsername: app.Config.AuthPassword,
	}

	promRouter.With(middleware.BasicAuth("Postmark Authentication", webhookCreds)).Post("/webhook", httputils.LoggingMiddleware(app.WebhookHandler))

	r.Mount("/", promRouter)
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
			defer func() {
				if err := file.Close(); err != nil {
					log.Println("Error closing file:", err)
				}
			}()
			if stat, err := file.Stat(); err == nil && !stat.IsDir() {
				http.ServeContent(w, r, r.URL.Path, stat.ModTime(), file.(io.ReadSeeker))
				return
			}
		}

		indexFile, err := distFS.Open("index.html")
		defer func() {
			if err := indexFile.Close(); err != nil {
				log.Println("Error closing index file:", err)
			}
		}()

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
