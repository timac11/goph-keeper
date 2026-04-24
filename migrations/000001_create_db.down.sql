DROP TRIGGER IF EXISTS update_user_updated_at ON "user";
DROP TRIGGER IF EXISTS update_secret_updated_at ON "secret";

DROP TABLE IF EXISTS "secret";
DROP TABLE IF EXISTS "user";

DROP FUNCTION IF EXISTS update_updated_at_column;
