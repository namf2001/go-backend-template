package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/namf2001/go-backend-template/config"
	authhandler "github.com/namf2001/go-backend-template/internal/handler/rest/v1/auth"
	usershandler "github.com/namf2001/go-backend-template/internal/handler/rest/v1/users"
)

// router defines the routes & handlers of the app
type router struct {
	ctx          context.Context
	cfg          *config.Config
	usersHandler *usershandler.Handler
	authHandler  *authhandler.Handler
}

// handler returns the handler for use by the server
func (rtr router) handler() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	rtr.routes(r)

	return r
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
	r.Group(rtr.apiV1)
}

func (rtr router) public(r chi.Router) {
	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

func (rtr router) apiV1(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", rtr.usersHandler.CreateUser)
			r.Get("/", rtr.usersHandler.ListUsers)
			r.Get("/{id}", rtr.usersHandler.GetUser)
			r.Put("/{id}", rtr.usersHandler.UpdateUser)
			r.Delete("/{id}", rtr.usersHandler.DeleteUser)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", rtr.authHandler.Login)
			r.Post("/register", rtr.authHandler.Register)
			r.Get("/google/login", rtr.authHandler.GoogleLogin)
			r.Get("/google/callback", rtr.authHandler.GoogleCallback)
		})
	})
}
