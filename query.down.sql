-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

DROP INDEX user_email_idx;
DROP TYPE "level";
DROP TABLE IF EXISTS "user" CASCADE;
DROP TABLE IF EXISTS "session" CASCADE;
DROP TABLE IF EXISTS "organization" CASCADE;
DROP TABLE IF EXISTS "box" CASCADE;
DROP TABLE IF EXISTS "permission" CASCADE;
