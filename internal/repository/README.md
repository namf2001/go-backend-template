# Repository Layer (`internal/repository/`)

Lớp này chịu trách nhiệm tương tác trực tiếp với cơ sở dữ liệu (Database) hoặc các nguồn dữ liệu bên ngoài.

## Trách nhiệm

-   Thực hiện các câu lệnh SQL (CRUD).
-   Mapping dữ liệu từ Database Row sang struct Model của Go.
-   Xử lý transaction.

## Nguyên tắc

-   Repository **chỉ** nên làm việc với Database (PostgreSQL, MySQL, Mongo...).
-   Repository **không** nên chứa business logic phức tạp.
-   Repository nên trả về Model và Error chuẩn của Go (tránh leak database driver error lên tầng trên).
