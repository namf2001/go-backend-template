# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Download dependencies first (cached)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
# CGO_ENABLED=0: Static build
# -ldflags="-w -s": Strip debug symbols (giảm size)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api ./cmd/server

# Run Stage
FROM alpine:3.21

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/api .

# Set timezone
ENV TZ=Asia/Ho_Chi_Minh

EXPOSE 8080

CMD ["./api"]
