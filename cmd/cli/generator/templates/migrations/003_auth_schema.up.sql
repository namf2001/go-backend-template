-- Alter users table to add new columns for Auth
-- We check if columns exist to avoid errors on repeated runs, though standard practice is clean migration.

ALTER TABLE users ADD COLUMN IF NOT EXISTS "emailVerified" TIMESTAMPTZ;
ALTER TABLE users ADD COLUMN IF NOT EXISTS image TEXT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS password VARCHAR(255);

-- Create accounts table for OAuth
CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  "userId" BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Changed to BIGINT to match users.id
  type VARCHAR(255) NOT NULL,
  provider VARCHAR(255) NOT NULL,
  "providerAccountId" VARCHAR(255) NOT NULL,
  refresh_token TEXT,
  access_token TEXT,
  expires_at BIGINT,
  id_token TEXT,
  scope TEXT,
  session_state TEXT,
  token_type TEXT
);

-- Create sessions table
CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL PRIMARY KEY,
  "userId" BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Changed to BIGINT to match users.id
  expires TIMESTAMPTZ NOT NULL,
  "sessionToken" VARCHAR(255) NOT NULL
);

-- Create verification_token table
CREATE TABLE IF NOT EXISTS verification_token (
  identifier TEXT NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  token TEXT NOT NULL,
  PRIMARY KEY (identifier, token)
);
