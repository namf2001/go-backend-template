# Database Migrations (`migrations/`)

Thư mục này chứa các file SQL để quản lý thay đổi schema của cơ sở dữ liệu (version control for database).

## Quy ước đặt tên

Các file migration nên tuân theo format:
`<timestamp/sequence>_<description>.<up/down>.sql`

Ví dụ:

-   `001_create_users_table.up.sql`: Script tạo bảng (được chạy khi `migrate-up`).
-   `001_create_users_table.down.sql`: Script xóa bảng (được chạy khi `migrate-down`).

## Cách sử dụng

Sử dụng `Makefile` để chạy migration:

-   `make migrate-up`: Chạy tất cả các file `.up.sql` chưa được apply.
-   `make migrate-down`: Chạy tất cả các file `.down.sql` để rollback.

## Lưu ý

-   **Không sửa đổi** file migration đã được merge/deploy. Nếu cần thay đổi, hãy tạo một migration mới.
-   Mỗi lần thay đổi schema (thêm/sửa/xóa bảng hoặc cột) đều cần tạo cặp file migration tương ứng.
