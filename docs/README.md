# Go Backend Template - Documentation Index

Chào mừng đến với Go Backend Template! Đây là hướng dẫn từng bước để tạo một backend API hoàn chỉnh theo Clean Architecture pattern.

---

## 📚 Tổng Quan

Project này cung cấp hướng dẫn chi tiết để xây dựng một REST API backend với Go, bao gồm:
- ✅ Clean Architecture pattern
- ✅ PostgreSQL database
- ✅ CRUD operations hoàn chỉnh
- ✅ Error handling và validation
- ✅ Sẵn sàng cho production

---

## 🚀 Bắt Đầu Nhanh

Nếu bạn muốn bắt đầu ngay, xem [Quick Start Guide](QUICKSTART.md).

---

## 📖 Hướng Dẫn Chi Tiết - 5 Phases

### [Phase 1: Project Setup & Dependencies](phases/phase1-setup.md)
**Thời gian:** ~30 phút

Trong phase này bạn sẽ:
- Khởi tạo Go module
- Tạo cấu trúc thư mục
- Cài đặt dependencies (Chi Router, PostgreSQL driver, v.v.)
- Setup Docker Compose cho PostgreSQL
- Tạo Makefile và configuration files

**Kết quả:** Project structure sẵn sàng, PostgreSQL chạy trong Docker

---

### [Phase 2: Database & Infrastructure](phases/phase2-infrastructure.md)
**Thời gian:** ~45 phút

Trong phase này bạn sẽ:
- Tạo database migrations
- Implement database connection pool
- Setup error handling system
- Tạo response utilities
- Implement configuration management

**Kết quả:** Infrastructure layer hoàn chỉnh, database sẵn sàng

---

### [Phase 3: Domain Models & Repository](phases/phase3-repository.md)
**Thời gian:** ~1 giờ

Trong phase này bạn sẽ:
- Tạo domain models (User)
- Implement repository interface
- Tạo PostgreSQL repository với raw SQL
- Implement validator utilities

**Kết quả:** Data access layer hoàn chỉnh

---

### [Phase 4: Controller & REST Handler](phases/phase4-controller-handler.md)
**Thời gian:** ~1.5 giờ

Trong phase này bạn sẽ:
- Implement business logic (Controller)
- Tạo REST API handlers
- Setup HTTP router với middleware
- Tạo main application
- Wire up tất cả components

**Kết quả:** Application hoàn chỉnh, có thể chạy được

---

### [Phase 5: Testing & Verification](phases/phase5-testing.md)
**Thời gian:** ~45 phút

Trong phase này bạn sẽ:
- Test tất cả API endpoints
- Verify error handling
- Test edge cases
- Tạo test scripts
- Verify database operations

**Kết quả:** Application đã được test đầy đủ và sẵn sàng sử dụng

---

## 🎯 Luồng Học Tập

### Cho Người Mới Bắt Đầu
1. Đọc [Quick Start Guide](QUICKSTART.md) để hiểu tổng quan
2. Làm theo từng phase theo thứ tự 1 → 5
3. Đọc kỹ giải thích trong mỗi phase
4. Chạy và test sau mỗi phase

### Cho Người Có Kinh Nghiệm
1. Review [Quick Start Guide](QUICKSTART.md)
2. Skim qua các phase để hiểu cấu trúc
3. Focus vào các phần quan tâm
4. Customize theo nhu cầu

---

## 📁 Cấu Trúc Project Sau Khi Hoàn Thành

```
go-backend-template/
├── cmd/
│   └── server/
│       ├── main.go              # Application entry point
│       └── router.go            # HTTP routes
├── internal/
│   ├── controller/
│   │   └── users/               # Business logic
│   ├── handler/
│   │   └── rest/v1/users/       # REST handlers
│   ├── repository/
│   │   └── users/               # Data access
│   ├── model/                   # Domain models
│   └── pkg/
│       ├── database/            # DB connection
│       ├── errors/              # Error types
│       ├── response/            # HTTP responses
│       └── validator/           # Input validation
├── migrations/                  # SQL migrations
├── config/                      # Configuration
├── docs/                        # Documentation
│   ├── phases/                  # Step-by-step guides
│   └── QUICKSTART.md
├── scripts/                     # Utility scripts
├── .env                         # Environment variables
├── docker-compose.yml           # PostgreSQL container
├── Makefile                     # Common commands
└── README.md
```

---

## 🛠️ Tech Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Language | Go 1.19+ | Backend programming |
| HTTP Router | Chi Router | HTTP routing & middleware |
| Database | PostgreSQL 15 | Data persistence |
| ORM | Raw SQL | Database queries |
| Validation | go-playground/validator | Input validation |
| Config | godotenv | Environment variables |
| Container | Docker | Development environment |

---

## 🎓 Những Gì Bạn Sẽ Học

### Architecture Patterns
- ✅ Clean Architecture
- ✅ Repository Pattern
- ✅ Dependency Injection
- ✅ Separation of Concerns

### Go Best Practices
- ✅ Project structure
- ✅ Error handling
- ✅ Interface design
- ✅ Context usage

### API Development
- ✅ REST API design
- ✅ HTTP middleware
- ✅ Request validation
- ✅ Error responses

### Database
- ✅ Migrations
- ✅ Connection pooling
- ✅ Raw SQL queries
- ✅ Transaction handling

---

## 📝 API Endpoints Overview

Sau khi hoàn thành, bạn sẽ có các endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/v1/users` | Create user |
| GET | `/api/v1/users` | List users (with pagination) |
| GET | `/api/v1/users/:id` | Get user by ID |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

---

## 🔧 Prerequisites

Trước khi bắt đầu, đảm bảo bạn đã cài đặt:

- **Go 1.19+**: [Download](https://go.dev/dl/)
- **Docker**: [Download](https://www.docker.com/products/docker-desktop)
- **Git**: [Download](https://git-scm.com/downloads)
- **PostgreSQL Client** (psql): Để test database
- **curl** hoặc **Postman**: Để test API

---

## ❓ Câu Hỏi Thường Gặp

### Tôi cần biết gì trước khi bắt đầu?
- Kiến thức cơ bản về Go
- Hiểu về HTTP và REST API
- Kiến thức cơ bản về SQL

### Mất bao lâu để hoàn thành?
- Toàn bộ: ~4-5 giờ
- Nếu đã có kinh nghiệm Go: ~2-3 giờ

### Tôi có thể customize không?
- Hoàn toàn có thể! Đây là template, bạn có thể thay đổi theo nhu cầu

### Có thể dùng cho production không?
- Có, nhưng nên thêm:
  - Authentication/Authorization
  - Logging và monitoring
  - Rate limiting
  - Unit tests
  - CI/CD pipeline

---

## 🚦 Bắt Đầu

Sẵn sàng? Bắt đầu với [Phase 1: Project Setup](phases/phase1-setup.md)!

Hoặc xem [Quick Start Guide](QUICKSTART.md) để có overview nhanh.

---

## 📞 Support

Nếu gặp vấn đề:
1. Kiểm tra lại từng bước trong phase tương ứng
2. Xem phần Troubleshooting trong Phase 5
3. Verify prerequisites đã được cài đặt đúng

---

## 🎉 Sau Khi Hoàn Thành

Sau khi hoàn thành tất cả 5 phases, bạn sẽ có:
- ✅ Một REST API hoàn chỉnh
- ✅ Hiểu rõ Clean Architecture
- ✅ Kinh nghiệm với Go best practices
- ✅ Foundation để build các project phức tạp hơn

**Next Steps:**
- Thêm authentication (JWT)
- Implement more domain models
- Add unit tests
- Deploy to cloud (AWS, GCP, Azure)
- Add API documentation (Swagger)
