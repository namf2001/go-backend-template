# Models (`internal/model/`)

Thư mục này chứa các định nghĩa cấu trúc dữ liệu (structs) được sử dụng trong toàn bộ ứng dụng.

## Nội dung

-   **Domain Models**: Các struct đại diện cho đối tượng nghiệp vụ (User, Product, Order...).
-   **Request/Response Models** (Optional): Các struct dùng riêng cho việc nhận request hoặc trả response nếu cần tách biệt với Domain Model.
-   **Service Interfaces**: Interface định nghĩa các hành vi (methods) mà Controller/Service cần implement (đôi khi được đặt ở `internal/core` trong Hexagonal Architecture).

## Nguyên tắc

-   Model nên "anemic" (chỉ chứa dữ liệu) hoặc "rich" (chứa validation logic cơ bản).
-   Tránh phụ thuộc vào các thư viện bên ngoài (external libs) nếu có thể.
