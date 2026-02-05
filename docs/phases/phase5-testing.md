# Phase 5: Testing & Verification

## Mục Tiêu
- Test API endpoints
- Verify toàn bộ flow hoạt động
- Tạo documentation và examples

---

## Bước 1: Start Application

```bash
# Đảm bảo PostgreSQL đang chạy
make docker-up

# Chờ database khởi động
sleep 5

# Run migrations
make migrate-up

# Start server
make run
```

Server sẽ chạy tại `http://localhost:8080`

---

## Bước 2: Test Health Check

```bash
curl http://localhost:8080/health
```

**Expected Response:**
```
OK
```

---

## Bước 3: Test Create User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "name": "John Doe"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "john@example.com",
    "name": "John Doe",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

---

## Bước 4: Test Get User

```bash
curl http://localhost:8080/api/v1/users/1
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "john@example.com",
    "name": "John Doe",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

---

## Bước 5: Test List Users

```bash
# List all users
curl http://localhost:8080/api/v1/users

# List with pagination
curl "http://localhost:8080/api/v1/users?limit=10&offset=0"

# Search by email
curl "http://localhost:8080/api/v1/users?email=john"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": 1,
        "email": "john@example.com",
        "name": "John Doe",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

---

## Bước 6: Test Update User

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith"
  }'
```

**Expected Response:**
```
HTTP 204 No Content
```

Verify update:
```bash
curl http://localhost:8080/api/v1/users/1
```

---

## Bước 7: Test Delete User

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

**Expected Response:**
```
HTTP 204 No Content
```

Verify deletion:
```bash
curl http://localhost:8080/api/v1/users/1
```

**Expected Response:**
```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "user not found"
  }
}
```

---

## Bước 8: Test Error Cases

### Duplicate Email
```bash
# Create first user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "name": "Test User"}'

# Try to create duplicate
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "name": "Another User"}'
```

**Expected Response:**
```json
{
  "success": false,
  "error": {
    "code": "ALREADY_EXISTS",
    "message": "user with this email already exists"
  }
}
```

### Invalid Email
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "invalid-email", "name": "Test"}'
```

**Expected Response:**
```json
{
  "success": false,
  "error": {
    "code": "INVALID_INPUT",
    "message": "validation failed"
  }
}
```

### Missing Fields
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com"}'
```

**Expected Response:**
```json
{
  "success": false,
  "error": {
    "code": "INVALID_INPUT",
    "message": "validation failed"
  }
}
```

---

## Bước 9: Verify Database

```bash
# Connect to database
psql -h localhost -U postgres -d go_backend_db

# Check users table
SELECT * FROM users;

# Check table structure
\d users

# Exit
\q
```

---

## Bước 10: Create Test Script

### File: `scripts/test-api.sh`

```bash
mkdir -p scripts

cat > scripts/test-api.sh << 'EOF'
#!/bin/bash

BASE_URL="http://localhost:8080"

echo "🧪 Testing Go Backend API"
echo "=========================="

# Health Check
echo -e "\n1. Health Check"
curl -s $BASE_URL/health
echo ""

# Create User
echo -e "\n2. Create User"
USER_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User"}')
echo $USER_RESPONSE | jq .

# Extract user ID
USER_ID=$(echo $USER_RESPONSE | jq -r '.data.id')

# Get User
echo -e "\n3. Get User (ID: $USER_ID)"
curl -s $BASE_URL/api/v1/users/$USER_ID | jq .

# List Users
echo -e "\n4. List Users"
curl -s $BASE_URL/api/v1/users | jq .

# Update User
echo -e "\n5. Update User"
curl -s -X PUT $BASE_URL/api/v1/users/$USER_ID \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name"}'
echo ""

# Verify Update
echo -e "\n6. Verify Update"
curl -s $BASE_URL/api/v1/users/$USER_ID | jq .

# Delete User
echo -e "\n7. Delete User"
curl -s -X DELETE $BASE_URL/api/v1/users/$USER_ID
echo ""

# Verify Deletion
echo -e "\n8. Verify Deletion (should return 404)"
curl -s $BASE_URL/api/v1/users/$USER_ID | jq .

echo -e "\n✅ All tests completed!"
EOF

chmod +x scripts/test-api.sh
```

Run test script:
```bash
./scripts/test-api.sh
```

---

## Bước 11: Create Postman Collection (Optional)

### File: `docs/postman-collection.json`

```bash
cat > docs/postman-collection.json << 'EOF'
{
  "info": {
    "name": "Go Backend Template API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "url": "http://localhost:8080/health"
      }
    },
    {
      "name": "Create User",
      "request": {
        "method": "POST",
        "header": [{"key": "Content-Type", "value": "application/json"}],
        "body": {
          "mode": "raw",
          "raw": "{\"email\":\"john@example.com\",\"name\":\"John Doe\"}"
        },
        "url": "http://localhost:8080/api/v1/users"
      }
    },
    {
      "name": "Get User",
      "request": {
        "method": "GET",
        "url": "http://localhost:8080/api/v1/users/1"
      }
    },
    {
      "name": "List Users",
      "request": {
        "method": "GET",
        "url": "http://localhost:8080/api/v1/users"
      }
    },
    {
      "name": "Update User",
      "request": {
        "method": "PUT",
        "header": [{"key": "Content-Type", "value": "application/json"}],
        "body": {
          "mode": "raw",
          "raw": "{\"name\":\"John Smith\"}"
        },
        "url": "http://localhost:8080/api/v1/users/1"
      }
    },
    {
      "name": "Delete User",
      "request": {
        "method": "DELETE",
        "url": "http://localhost:8080/api/v1/users/1"
      }
    }
  ]
}
EOF
```

---

## Kết Quả Mong Đợi

✅ Tất cả API endpoints hoạt động đúng  
✅ Error handling hoạt động chính xác  
✅ Database operations thành công  
✅ Validation rules được enforce  
✅ Test script chạy thành công  

---

## Troubleshooting

### Server không start được
```bash
# Check port đã được sử dụng chưa
lsof -i :8080

# Kill process nếu cần
kill -9 <PID>
```

### Database connection failed
```bash
# Check PostgreSQL đang chạy
docker ps | grep postgres

# Restart PostgreSQL
make docker-down
make docker-up
```

### Migration failed
```bash
# Drop và recreate database
psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS go_backend_db;"
psql -h localhost -U postgres -c "CREATE DATABASE go_backend_db;"

# Run migration lại
make migrate-up
```

---

## Hoàn Thành! 🎉

Bạn đã hoàn thành việc tạo một Go backend project hoàn chỉnh với:

✅ Clean Architecture pattern  
✅ REST API với CRUD operations  
✅ Database migrations  
✅ Error handling  
✅ Input validation  
✅ Comprehensive testing  

---

## Next Steps

1. **Add Authentication**: Implement JWT authentication
2. **Add Logging**: Structured logging với Zap/Zerolog
3. **Add Metrics**: Prometheus metrics
4. **Add Tests**: Unit tests và integration tests
5. **Add Documentation**: Swagger/OpenAPI docs
6. **Dockerize**: Create Dockerfile cho deployment
