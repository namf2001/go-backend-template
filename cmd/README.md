# Command Line Applications (`cmd/`)

Thư mục này chứa các điểm khởi chạy (entry points) của ứng dụng. Mỗi thư mục con ở đây tương ứng với một ứng dụng thực thi được (executable).

## Cấu trúc

-   `server/`: Chứa hàm `main` để khởi chạy HTTP API server.
-   `jobs/`: Chứa các background jobs hoặc worker processes (nếu có).

## Nguyên tắc

-   Tuyệt đối **không** đặt business logic ở đây.
-   Hàm `main` chỉ nên làm nhiệm vụ:
    -   Load cấu hình (config).
    -   Khởi tạo các dependencies (database, loggers, services).
    -   Chạy ứng dụng (vd: `server.Run()`).
