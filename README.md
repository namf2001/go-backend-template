# Go Backend Template - Clean Architecture

Đây là một template project backend Go theo Clean Architecture pattern, tương tự như Thor API.

## 📚 Cấu Trúc Project

```
go-backend-template/
├── cmd/                    # Điểm khởi chạy của ứng dụng (Entry points)
│   ├── server/             # Chứa hàm main để chạy HTTP server
│   └── jobs/               # Chứa các background jobs (nếu có)
├── config/                 # Chứa cấu hình của ứng dụng (load env, config struct)
├── docs/                   # Tài liệu của dự án
│   └── phases/             # Hướng dẫn từng bước phát triển dự án
├── internal/               # Mã nguồn nội bộ của ứng dụng (Private application code)
│   ├── controller/         # Xử lý logic nghiệp vụ (Business Logic Layer)
│   ├── handler/            # Xử lý HTTP request/response (Transport Layer)
│   │   ├── rest/           # RESTful API handlers
│   │   └── middleware/     # HTTP Middlewares (Auth, Logging, etc.)
│   ├── model/              # Định nghĩa các struct dữ liệu (Data Models)
│   ├── repository/         # Tương tác với cơ sở dữ liệu (Data Access Layer)
│   └── pkg/                # Các gói tiện ích nội bộ (Internal shared packages)
│       ├── database/       # Kết nối DB
│       ├── errors/         # Định nghĩa lỗi chung
│       ├── jwt/            # Xử lý JWT Token
│       ├── response/       # Chuẩn hóa format response
│       └── validator/      # Validate dữ liệu đầu vào
├── migrations/             # Chứa các file migration SQL (.up.sql, .down.sql)
├── Makefile                # Các lệnh automation (build, run, migrate...)
├── docker-compose.yml      # Cấu hình Docker cho development (DB, Redis...)
├── go.mod                  # Quản lý dependencies
└── README.md               # Tài liệu chính của dự án
```

## 🚀 Quick Start

Làm theo các phase trong thư mục `docs/phases/`:

1. **Phase 1**: Project Setup & Dependencies
2. **Phase 2**: Database & Infrastructure
3. **Phase 3**: Core Application Layer
4. **Phase 4**: Sample Flow (User Management)
5. **Phase 5**: Testing & Verification

## 📖 Chi Tiết

Xem file trong `docs/phases/` để có hướng dẫn chi tiết từng bước với:
- Các câu lệnh cần chạy
- Code cần tạo
- Giải thích từng phần

## 🎯 Mục Tiêu

Sau khi hoàn thành, bạn sẽ có:
- ✅ REST API hoàn chỉnh với CRUD operations
- ✅ Clean Architecture với separation of concerns
- ✅ Database migrations
- ✅ Error handling và validation
- ✅ Sẵn sàng để mở rộng

## 🛠️ Tech Stack

- **Go 1.19+**
- **Chi Router** - HTTP routing
- **PostgreSQL** - Database
- **SQLBoiler** - ORM
- **Docker** - Development environment
