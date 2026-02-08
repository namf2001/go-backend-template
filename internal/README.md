# Internal Application Code (`internal/`)

Đây là nơi chứa toàn bộ mã nguồn nghiệp vụ của dự án. Theo quy ước chuẩn của Go, code trong thư mục `internal/` **không thể** được import bởi các project khác bên ngoài module này.

## Cấu trúc

-   `controller/`: Business Logic Layer - Xử lý nghiệp vụ chính.
-   `handler/`: Transport Layer - Xử lý request/response (HTTP, RPC).
-   `model/`: Data Models - Định nghĩa các struct dữ liệu.
-   `repository/`: Data Access Layer - Tương tác trực tiếp với Database.
-   `pkg/`: Internal Shared Packages - Các thư viện tiện ích dùng chung nội bộ.

## Luồng dữ liệu (Data Flow)

User Request -> `handler` -> `controller` -> `repository` -> Database
