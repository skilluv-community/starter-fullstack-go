package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/skilluv-community/starter-fullstack-go/backend/internal/db"
	"github.com/skilluv-community/starter-fullstack-go/backend/internal/routes"
)

type Config struct {
	DatabaseURL    string `env:"DATABASE_URL,required"`
	BackendPort    string `env:"BACKEND_PORT" envDefault:"3001"`
	CORSOrigin     string `env:"CORS_ORIGIN" envDefault:"http://localhost:5173"`
	MigrationsPath string `env:"MIGRATIONS_PATH" envDefault:"migrations"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		slog.Error("config", "err", err)
		os.Exit(1)
	}

	if err := db.Migrate(cfg.DatabaseURL, cfg.MigrationsPath); err != nil {
		slog.Error("migrate", "err", err)
		os.Exit(1)
	}

	pool, err := db.Open(cfg.DatabaseURL)
	if err != nil {
		slog.Error("db.open", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(cfg.CORSOrigin, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	r.GET("/health", routes.Health)

	api := r.Group("/api")
	api.GET("/hello", routes.Hello)

	nh := routes.NotesHandler{DB: pool}
	api.GET("/notes", nh.List)
	api.POST("/notes", nh.Create)
	api.DELETE("/notes/:id", nh.Delete)

	slog.Info("listening", "addr", "0.0.0.0:"+cfg.BackendPort)
	if err := r.Run("0.0.0.0:" + cfg.BackendPort); err != nil {
		slog.Error("server", "err", err)
		os.Exit(1)
	}
}
