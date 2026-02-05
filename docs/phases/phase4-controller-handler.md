# Phase 4: Controller & REST Handler

## Mục Tiêu
- Implement business logic layer (Controller)
- Tạo REST API handlers
- Setup HTTP router và middleware
- Tạo main application

---

## Bước 1: Implement Controller Interface

### File: `internal/controller/users/controller.go`

```go
cat > internal/controller/users/controller.go << 'EOF'
package users

import (
	"context"

	"github.com/yourusername/go-backend-template/internal/model"
)

// Controller defines the business logic interface for users
type Controller interface {
	CreateUser(ctx context.Context, input CreateUserInput) (UserOutput, error)
	GetUser(ctx context.Context, id int64) (UserOutput, error)
	ListUsers(ctx context.Context, filters ListFilters) (ListUsersOutput, error)
	UpdateUser(ctx context.Context, id int64, input UpdateUserInput) error
	DeleteUser(ctx context.Context, id int64) error
}

// CreateUserInput represents input for creating a user
type CreateUserInput struct {
	Email string `validate:"required,email"`
	Name  string `validate:"required,min=2,max=100"`
}

// UpdateUserInput represents input for updating a user
type UpdateUserInput struct {
	Email string `validate:"omitempty,email"`
	Name  string `validate:"omitempty,min=2,max=100"`
}

// UserOutput represents user output
type UserOutput struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ListFilters represents filters for listing users
type ListFilters struct {
	Limit  int
	Offset int
	Email  string
}

// ListUsersOutput represents output for listing users
type ListUsersOutput struct {
	Users []UserOutput `json:"users"`
	Total int          `json:"total"`
}

// ToOutput converts model.User to UserOutput
func ToOutput(user model.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
EOF
```

---

## Bước 2: Implement Controller Logic

### File: `internal/controller/users/impl.go`

```go
cat > internal/controller/users/impl.go << 'EOF'
package users

import (
	"context"

	"github.com/pkg/errors"
	"github.com/yourusername/go-backend-template/internal/model"
	"github.com/yourusername/go-backend-template/internal/pkg/validator"
	"github.com/yourusername/go-backend-template/internal/repository/users"
	apperrors "github.com/yourusername/go-backend-template/internal/pkg/errors"
)

type impl struct {
	repo users.Repository
}

// New creates a new users controller
func New(repo users.Repository) Controller {
	return &impl{repo: repo}
}

// CreateUser creates a new user
func (i *impl) CreateUser(ctx context.Context, input CreateUserInput) (UserOutput, error) {
	// Validate input
	if err := validator.Validate(input); err != nil {
		return UserOutput{}, apperrors.InvalidInput("validation failed")
	}

	// Check if user already exists
	_, err := i.repo.GetByEmail(ctx, input.Email)
	if err == nil {
		return UserOutput{}, apperrors.AlreadyExists("user with this email already exists")
	}

	// Create user
	user := model.User{
		Email: input.Email,
		Name:  input.Name,
	}

	created, err := i.repo.Create(ctx, user)
	if err != nil {
		return UserOutput{}, errors.Wrap(err, "failed to create user")
	}

	return ToOutput(created), nil
}

// GetUser retrieves a user by ID
func (i *impl) GetUser(ctx context.Context, id int64) (UserOutput, error) {
	user, err := i.repo.GetByID(ctx, id)
	if err != nil {
		return UserOutput{}, err
	}

	return ToOutput(user), nil
}

// ListUsers retrieves users with filters
func (i *impl) ListUsers(ctx context.Context, filters ListFilters) (ListUsersOutput, error) {
	repoFilters := users.ListFilters{
		Limit:  filters.Limit,
		Offset: filters.Offset,
		Email:  filters.Email,
	}

	userList, err := i.repo.List(ctx, repoFilters)
	if err != nil {
		return ListUsersOutput{}, errors.Wrap(err, "failed to list users")
	}

	outputs := make([]UserOutput, len(userList))
	for idx, user := range userList {
		outputs[idx] = ToOutput(user)
	}

	return ListUsersOutput{
		Users: outputs,
		Total: len(outputs),
	}, nil
}

// UpdateUser updates an existing user
func (i *impl) UpdateUser(ctx context.Context, id int64, input UpdateUserInput) error {
	// Validate input
	if err := validator.Validate(input); err != nil {
		return apperrors.InvalidInput("validation failed")
	}

	// Get existing user
	user, err := i.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Update fields
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Name != "" {
		user.Name = input.Name
	}

	// Save changes
	if err := i.repo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

// DeleteUser deletes a user
func (i *impl) DeleteUser(ctx context.Context, id int64) error {
	if err := i.repo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}
EOF
```

---

## Bước 3: Implement REST Handler

### File: `internal/handler/rest/v1/users/handler.go`

```go
cat > internal/handler/rest/v1/users/handler.go << 'EOF'
package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yourusername/go-backend-template/internal/controller/users"
	"github.com/yourusername/go-backend-template/internal/pkg/response"
)

type Handler struct {
	controller users.Controller
}

// New creates a new users handler
func New(controller users.Controller) *Handler {
	return &Handler{controller: controller}
}

// CreateUser handles POST /users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input users.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.controller.CreateUser(r.Context(), input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, user)
}

// GetUser handles GET /users/:id
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.controller.GetUser(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}

// ListUsers handles GET /users
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	email := r.URL.Query().Get("email")

	limit := 10 // default
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	offset := 0
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	filters := users.ListFilters{
		Limit:  limit,
		Offset: offset,
		Email:  email,
	}

	result, err := h.controller.ListUsers(r.Context(), filters)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}

// UpdateUser handles PUT /users/:id
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	var input users.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, err)
		return
	}

	if err := h.controller.UpdateUser(r.Context(), id, input); err != nil {
		response.Error(w, err)
		return
	}

	response.NoContent(w)
}

// DeleteUser handles DELETE /users/:id
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := h.controller.DeleteUser(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}

	response.NoContent(w)
}
EOF
```

---

## Bước 4: Setup Router

### File: `cmd/server/router.go`

```go
cat > cmd/server/router.go << 'EOF'
package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	usershandler "github.com/yourusername/go-backend-template/internal/handler/rest/v1/users"
)

func setupRouter(usersHandler *usershandler.Handler) http.Handler {
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

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Users routes
		r.Post("/users", usersHandler.CreateUser)
		r.Get("/users", usersHandler.ListUsers)
		r.Get("/users/{id}", usersHandler.GetUser)
		r.Put("/users/{id}", usersHandler.UpdateUser)
		r.Delete("/users/{id}", usersHandler.DeleteUser)
	})

	return r
}
EOF
```

---

## Bước 5: Create Main Application

### File: `cmd/server/main.go`

```go
cat > cmd/server/main.go << 'EOF'
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/go-backend-template/config"
	userscontroller "github.com/yourusername/go-backend-template/internal/controller/users"
	usershandler "github.com/yourusername/go-backend-template/internal/handler/rest/v1/users"
	"github.com/yourusername/go-backend-template/internal/pkg/database"
	usersrepo "github.com/yourusername/go-backend-template/internal/repository/users"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

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
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("✓ Database connected successfully")

	// Initialize repositories
	usersRepository := usersrepo.NewPostgresRepository(db)

	// Initialize controllers
	usersController := userscontroller.New(usersRepository)

	// Initialize handlers
	usersHandler := usershandler.New(usersController)

	// Setup router
	router := setupRouter(usersHandler)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 Environment: %s", cfg.Server.Env)
	log.Printf("🔗 Health check: http://localhost%s/health", addr)
	log.Printf("🔗 API base URL: http://localhost%s/api/v1", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
EOF
```

---

## Bước 6: Update Go Module Path

Cập nhật tất cả import paths trong các file vừa tạo:

```bash
# Thay thế "github.com/yourusername/go-backend-template" 
# bằng module path thực tế từ go.mod

# Ví dụ nếu module của bạn là "myapp"
find . -type f -name "*.go" -exec sed -i '' 's|github.com/yourusername/go-backend-template|myapp|g' {} +
```

---

## Kiểm Tra

```bash
# Tidy dependencies
go mod tidy

# Build application
make build

# Run application
make run
```

---

## Kết Quả Mong Đợi

✅ Controller layer đã được implement  
✅ REST handlers đã sẵn sàng  
✅ Router với middleware đã được setup  
✅ Main application có thể chạy  

---

## Tiếp Theo

Chuyển sang **Phase 5: Testing & Verification** để:
- Test API endpoints
- Viết documentation
- Tạo examples
