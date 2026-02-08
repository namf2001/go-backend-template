# Controller Layer (`internal/controller/`)

Đây là nơi chứa toàn bộ logic nghiệp vụ (Business Logic) của ứng dụng.

## Trách nhiệm

-   Nhận dữ liệu đã được xử lý sơ bộ từ `handler`.
-   Thực hiện các quy tắc nghiệp vụ (Business Rules).
-   Gọi `repository` để truy xuất/lưu trữ dữ liệu.
-   Trả kết quả về cho `handler`.

## Nguyên tắc

-   Controller không nên biết về HTTP (không nên có `w http.ResponseWriter` hay `r *http.Request` ở đây).
-   Controller chỉ nên làm việc với các Model và Error thuần túy của Go.
-   Việc tách biệt này giúp Controller dễ dàng được unit test và tái sử dụng (ví dụ: gọi từ gRPC handler hoặc Background Job).
