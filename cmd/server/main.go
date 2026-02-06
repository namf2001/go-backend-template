package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namf2001/go-backend-template/config"
	userscontroller "github.com/namf2001/go-backend-template/internal/controller/users"
	usershandler "github.com/namf2001/go-backend-template/internal/handler/rest/v1/users"
	"github.com/namf2001/go-backend-template/internal/pkg/database"
	"github.com/namf2001/go-backend-template/internal/repository"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := run(ctx, cfg); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	// Connect to database
	dbCfg := database.Config{
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
		User:         cfg.Database.User,
		Password:     cfg.Database.Password,
		DBName:       cfg.Database.DBName,
		SSLMode:      cfg.Database.SSLMode,
		MaxOpenConns: cfg.Database.MaxOpenConns,
		MaxIdleConns: cfg.Database.MaxIdleConns,
	}

	db, err := database.NewPostgresConnection(dbCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	log.Println("✓ Database connected successfully")

	// Initialize repository
	repo := repository.New(db)

	// Initialize controllers
	usersController := userscontroller.New(repo)

	// Initialize handlers
	usersHandler := usershandler.New(usersController)

	// Setup router
	rtr := router{
		ctx:          ctx,
		cfg:          cfg,
		usersHandler: usersHandler,
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      rtr.handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 Environment: %s", cfg.Server.Env)
	log.Printf("🔗 Health check: http://localhost%s/health", addr)
	log.Printf("🔗 API base URL: http://localhost%s/api/v1/users", addr)

	// Graceful shutdown channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-done
	log.Println("Server stopping...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	log.Println("Server exited properly")
	return nil
}
