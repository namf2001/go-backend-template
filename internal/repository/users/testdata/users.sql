-- Test data for users repository tests
-- This file is loaded by testdb.LoadTestSQLFile within a rolled-back transaction

DELETE FROM users;

INSERT INTO users (id, email, name, password, image, "emailVerified", created_at, updated_at)
VALUES
    (1001, 'test1@example.com', 'Test User 1', '$2a$10$hashedpassword1', 'https://example.com/img1.png', '2024-01-01 00:00:00+07', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
    (1002, 'test2@example.com', 'Test User 2', '$2a$10$hashedpassword2', '', NULL, '2024-01-02 00:00:00', '2024-01-02 00:00:00'),
    (1003, 'admin@example.com', 'Admin User', '$2a$10$hashedpassword3', 'https://example.com/admin.png', '2024-01-03 00:00:00+07', '2024-01-03 00:00:00', '2024-01-03 00:00:00');
