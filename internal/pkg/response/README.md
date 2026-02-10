# Response Package

This package provides standardized JSON response utilities for the Go backend.

## Usage

### Success Responses

Use `response.Success` for standard 200 OK responses with data:

```go
response.Success(w, data)
```

Use `response.Created` for 201 Created responses:

```go
response.Created(w, data)
```

Use `response.NoContent` for 204 No Content responses:

```go
response.NoContent(w)
```

### Error Responses

Use `response.Error` to handle errors. It supports both `*apperrors.AppError` and standard sentinel errors from the `internal/pkg/errors` package.

```go
response.Error(w, err)
```

**How it works:**

1.  **AppErrors**: If the error is an `*apperrors.AppError`, it maps the error code to the appropriate HTTP status code.
2.  **Sentinel Errors**: If it's a standard error, it checks if it matches known sentinels (e.g., `apperrors.ErrNotFound`).
3.  **Fallback**: If the error is unknown, it returns a 500 Internal Server Error.

### Error Mapping

| Error Code / Sentinel | HTTP Status |
| :--- | :--- |
| `NOT_FOUND` / `ErrNotFound` | 404 Not Found |
| `ALREADY_EXISTS` / `ErrAlreadyExists` | 409 Conflict |
| `INVALID_INPUT` / `ErrInvalidInput` | 400 Bad Request |
| `UNAUTHORIZED` / `ErrUnauthorized` | 401 Unauthorized |
| `FORBIDDEN` / `ErrForbidden` | 403 Forbidden |
| `CONFLICT` / `ErrConflict` | 409 Conflict |
| `INTERNAL_ERROR` / Other | 500 Internal Server Error |
