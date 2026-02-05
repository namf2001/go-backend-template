# Quick Start Guide

Hướng dẫn nhanh để bắt đầu với Go Backend Template.

---

## Prerequisites

- Go 1.19 hoặc cao hơn
- Docker và Docker Compose
- PostgreSQL client tools (psql)
- curl hoặc Postman (để test API)

---

## Installation

### 1. Clone hoặc Navigate to Project

```bash
cd /Users/flnam2001/Workplace/PersonalProjects/Backend/go-backend-template
```

### 2. Làm theo từng Phase

Project này được tổ chức thành 5 phases, mỗi phase có hướng dẫn chi tiết trong `docs/phases/`:

#### **Phase 1: Setup** (`docs/phases/phase1-setup.md`)
- Khởi tạo Go module
- Tạo cấu trúc thư mục
- Cài đặt dependencies
- Setup Docker Compose

#### **Phase 2: Infrastructure** (`docs/phases/phase2-infrastructure.md`)
- Database migrations
- Database connection
- Error handling
- Response utilities

#### **Phase 3: Repository** (`docs/phases/phase3-repository.md`)
- Domain models
- Repository pattern
- Database queries

#### **Phase 4: Controller & Handler** (`docs/phases/phase4-controller-handler.md`)
- Business logic (Controller)
- REST API handlers
- HTTP router
- Main application

#### **Phase 5: Testing** (`docs/phases/phase5-testing.md`)
- API testing
- Verification
- Test scripts

---

## Quick Commands

```bash
# Start PostgreSQL
make docker-up

# Run migrations
make migrate-up

# Run server
make run

# Run tests
make test

# Build binary
make build

# Stop PostgreSQL
make docker-down
```

---

## API Endpoints

Base URL: `http://localhost:8080/api/v1`

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/users` | Create user |
| GET | `/users` | List users |
| GET | `/users/:id` | Get user by ID |
| PUT | `/users/:id` | Update user |
| DELETE | `/users/:id` | Delete user |

---

## Example Usage

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","name":"John Doe"}'
```

### Get User
```bash
curl http://localhost:8080/api/v1/users/1
```

### List Users
```bash
curl http://localhost:8080/api/v1/users
```

---

## Project Structure

```
go-backend-template/
├── cmd/server/              # Application entry point
├── internal/
│   ├── controller/          # Business logic
│   ├── handler/             # HTTP handlers
│   ├── repository/          # Data access
│   ├── model/               # Domain models
│   └── pkg/                 # Utilities
├── migrations/              # Database migrations
├── config/                  # Configuration
└── docs/                    # Documentation
    └── phases/              # Step-by-step guides
```

---

## Architecture

Project này sử dụng **Clean Architecture** với các layer:

```
Handler (REST API)
    ↓
Controller (Business Logic)
    ↓
Repository (Data Access)
    ↓
Database (PostgreSQL)
```

---

## Tech Stack

- **Go 1.19+** - Programming language
- **Chi Router** - HTTP routing
- **PostgreSQL** - Database
- **Docker** - Containerization
- **Raw SQL** - Database queries

---

## Support

Xem chi tiết trong các file phase tương ứng:
- `docs/phases/phase1-setup.md`
- `docs/phases/phase2-infrastructure.md`
- `docs/phases/phase3-repository.md`
- `docs/phases/phase4-controller-handler.md`
- `docs/phases/phase5-testing.md`
