package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib"
	"nusantarakita/internal/handler"
	"nusantarakita/internal/repository"
	"nusantarakita/internal/service"
)

func main() {
	// Set default structured JSON logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.Error("DATABASE_URL environment variable is required")
		os.Exit(1)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	// db.Close() will be handled during graceful shutdown below

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	// Initialize layers
	repo := repository.NewDaerahRepository(db)
	svc := service.NewDaerahService(repo)
	h := handler.New(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	// Simple CORS middleware
	corsHandler := corsMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler,
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		slog.Info("Server listening", "port", port)
		serverErrors <- srv.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	case sig := <-shutdown:
		slog.Info("Start shutdown...", "signal", sig.String())

		// Create context with timeout for the shutdown process
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Shutdown the server gracefully
		if err := srv.Shutdown(ctx); err != nil {
			slog.Warn("Graceful shutdown did not complete in 10s", "error", err)
			if err := srv.Close(); err != nil {
				slog.Error("Could not stop server gracefully", "error", err)
				os.Exit(1)
			}
		}

		// Close database connection
		if err := db.Close(); err != nil {
			slog.Warn("Could not close database connection gracefully", "error", err)
		}

		slog.Info("Graceful shutdown complete")
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
