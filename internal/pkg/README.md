# Internal Packages (`internal/pkg/`)

Thư mục này chứa các thư viện tiện ích (utility helpers) được sử dụng chung trong nội bộ `internal/` của dự án. Không nên expose các package này ra ngoài (nếu muốn reusable cho nhiều project khác nhau, hãy đưa vào `pkg/` ở root level).

## Các package phổ biến

-   `database/`: Khởi tạo kết nối DB logic.
-   `response/`: Các hàm helper để trả về JSON response chuẩn (Success, Error).
-   `errors/`: Định nghĩa các custom application errors.
-   `logger/`: Cấu hình logging.
-   `validator/`: Helper validate struct.
