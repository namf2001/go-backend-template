# Configuration (`config/`)

Thư mục này chịu trách nhiệm load và quản lý cấu hình của toàn bộ ứng dụng.

## Chức năng

-   Định nghĩa các struct chứa config (vd: `DatabaseConfig`, `ServerConfig`).
-   Load cấu hình từ biến môi trường (Environment Variables) hoặc file `.env`.
-   Cung cấp các giá trị mặc định (defaults) nếu thiếu config.

## Nguyên tắc

-   Sử dụng thư viện như `github.com/kelseyhightower/envconfig` hoặc `github.com/spf13/viper` để load config.
-   Không hardcode các giá trị nhạy cảm (secrets, passwords) trong code.
