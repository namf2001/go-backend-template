# Phase 1: Project Setup & Dependencies

## Mục Tiêu
- Khởi tạo Go module
- Tạo cấu trúc thư mục
- Cài đặt dependencies cơ bản

---

## Bước 1: Khởi Tạo Go Module

```bash
# Di chuyển vào thư mục project
cd /Users/flnam2001/Workplace/PersonalProjects/Backend/go-backend-template

# Khởi tạo Go module
go mod init github.com/yourusername/go-backend-template
```

---

## Bước 2: Tạo Cấu Trúc Thư Mục

```bash
# Tạo các thư mục chính
mkdir -p cmd/server
mkdir -p internal/{controller,handler,repository,model,pkg}
mkdir -p internal/handler/rest/v1/users
mkdir -p internal/controller/users
mkdir -p internal/repository/users
mkdir -p internal/pkg/{database,errors,validator,response}
mkdir -p migrations
mkdir -p config
```

---

## Bước 3: Cài Đặt Dependencies

```bash
# HTTP Router
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors

# Database
go get github.com/lib/pq
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/null/v8

# Utilities
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/pkg/errors
go get github.com/google/uuid
```

---

## Bước 4: Tạo File Cấu Hình Cơ Bản

### File: `.env.example`

```bash
cat > .env.example << 'EOF'
# Server Configuration
SERVER_PORT=8080
SERVER_ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_backend_db
DB_SSL_MODE=disable

# Database Pool
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
EOF
```

### File: `.env`

```bash
# Copy .env.example to .env
cp .env.example .env
```

---

## Bước 5: Tạo Docker Compose cho PostgreSQL

### File: `docker-compose.yml`

```bash
cat > docker-compose.yml << 'EOF'
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: go-backend-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: go_backend_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

volumes:
  postgres_data:
EOF
```

---

## Bước 6: Tạo Makefile

### File: `Makefile`

```bash
cat > Makefile << 'EOF'
.PHONY: help run build test clean docker-up docker-down migrate-up migrate-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	go run cmd/server/main.go

build: ## Build the application
	go build -o bin/server cmd/server/main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build artifacts
	rm -rf bin/

docker-up: ## Start Docker containers
	docker-compose up -d

docker-down: ## Stop Docker containers
	docker-compose down

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	psql -h localhost -U postgres -d go_backend_db -f migrations/001_create_users_table.sql

migrate-down: ## Rollback database migrations
	@echo "Rolling back migrations..."
	psql -h localhost -U postgres -d go_backend_db -c "DROP TABLE IF EXISTS users;"

deps: ## Download dependencies
	go mod download
	go mod tidy

.DEFAULT_GOAL := help
EOF
```

---

## Bước 7: Tạo .gitignore

```bash
cat > .gitignore << 'EOF'
# Binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# Environment variables
.env

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
EOF
```

---

## Kiểm Tra

Sau khi hoàn thành Phase 1, chạy các lệnh sau để kiểm tra:

```bash
# Kiểm tra Go module
go mod tidy

# Kiểm tra cấu trúc thư mục
tree -L 3 -I 'vendor'

# Start PostgreSQL
make docker-up

# Kiểm tra PostgreSQL đã chạy
docker ps | grep postgres
```

---

## Kết Quả Mong Đợi

✅ Go module đã được khởi tạo  
✅ Cấu trúc thư mục đã được tạo  
✅ Dependencies đã được cài đặt  
✅ PostgreSQL đang chạy trong Docker  

---

## Tiếp Theo

Chuyển sang **Phase 2: Database & Infrastructure** để:
- Tạo database migrations
- Setup database connection
- Implement error handling
