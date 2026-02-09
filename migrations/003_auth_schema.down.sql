-- Drop tables
DROP TABLE IF EXISTS verification_token;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS accounts;

-- Revert users table changes (optional, usually dropping columns is risky but for down migration it is expected)
ALTER TABLE users DROP COLUMN IF EXISTS password;
ALTER TABLE users DROP COLUMN IF EXISTS image;
ALTER TABLE users DROP COLUMN IF EXISTS "emailVerified";
