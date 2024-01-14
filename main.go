package main

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"osr/db"
	"osr/views"

	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

//go:embed assets
var assets embed.FS

//go:embed db/migrations/*.sql
var migrations embed.FS

//go:embed feeds.txt
var feedFile []byte

func FeedUrls() []string {
	var feedUrls []string
	scanner := bufio.NewScanner(bytes.NewReader(feedFile))
	for scanner.Scan() {
		feedUrls = append(feedUrls, scanner.Text())
	}
	return feedUrls
}

func main() {
	if err := Migrate(os.Getenv("DATABASE_URL")); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)
	Feeds(q)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		items, err := q.GetItems(req.Context())
		if err != nil {
			slog.Error(err.Error())
			// TODO
		}
		if err := views.Home(items).Render(req.Context(), w); err != nil {
			slog.Error(err.Error())
		}
	})
	r.Handle("/assets/*", http.FileServer(http.FS(assets)))

	slog.Info("listening", "addr", ":3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		slog.Error(err.Error())
	}
}
