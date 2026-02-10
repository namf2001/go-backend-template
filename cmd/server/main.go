package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namf2001/go-backend-template/config"
	authcontroller "github.com/namf2001/go-backend-template/internal/controller/auth"
	userscontroller "github.com/namf2001/go-backend-template/internal/controller/users"
	authhandler "github.com/namf2001/go-backend-template/internal/handler/rest/v1/auth"
	usershandler "github.com/namf2001/go-backend-template/internal/handler/rest/v1/users"
	"github.com/namf2001/go-backend-template/internal/pkg/database"
	"github.com/namf2001/go-backend-template/internal/pkg/oauth"
	"github.com/namf2001/go-backend-template/internal/repository"
)

// @title           Go Backend Template API
// @version         1.0
// @description     This is a sample server implementation.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()

	// Parse command line flags for environment; allow APP_ENV fallback
	envFlag := flag.String("e", "", "Environment to run (development, staging, production)")
	flag.Parse()

	env := *envFlag
	if env == "" {
		env = os.Getenv("APP_ENV")
		if env == "" {
			env = "dev"
		}
	}

	log.Printf("Initializing config for environment: %s", env)
	config.Init(env)

	if err := run(ctx); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg := config.GetConfig()

	db, err := database.NewPostgresConnection()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	log.Println("✓ Database connected successfully")

	// Initialize OAuth
	oauth.Init()

	// Initialize repository
	repo := repository.New(db)

	// Initialize controllers
	usersController := userscontroller.New(repo)
	authController := authcontroller.New(repo)

	// Initialize handlers
	usersHandler := usershandler.New(usersController)
	authHandler := authhandler.New(authController)

	// Setup router
	rtr := router{
		ctx:          ctx,
		usersHandler: usersHandler,
		authHandler:  authHandler,
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.GetString("APP_PORT"))
	srv := &http.Server{
		Addr:         addr,
		Handler:      rtr.handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 Environment: %s", cfg.GetString("APP_ENV"))
	log.Printf("🔗 Health check: http://localhost%s/health", addr)
	log.Printf("🔗 API Swagger URL: http://localhost%s/swagger/index.html", addr)
	log.Printf("🔗 API Metrics URL: http://localhost%s/metrics", addr)

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
