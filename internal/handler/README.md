# Handler Layer (`internal/handler/`)

Đây là lớp tiếp nhận request từ bên ngoài (thường là HTTP request), parse, validate cơ bản và chuyển tiếp đến Controller (Business Logic Layer).

## Trách nhiệm

-   Parse payload từ request body, query params, path params.
-   Validate dữ liệu đầu vào cơ bản (required fields, format...).
-   Gọi chức năng tương ứng ở tầng `controller`.
-   Format kết quả trả về cho client (JSON response, Status code...).

## Cấu trúc

-   `rest/`: Chứa các handler cho RESTful API.
-   `middleware/`: Chứa các middleware xử lý trước/sau request (Auth, Logger, Recovery...).
