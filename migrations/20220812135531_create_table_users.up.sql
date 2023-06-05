-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE "level" AS ENUM ('user', 'admin', 'integration');

CREATE TABLE users (
	id			     						UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
	first_name   						VARCHAR(20)                 NOT NULL CHECK ( first_name <> '' ),
	last_name    						VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
	email        						VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
	"password"     					VARCHAR(150)                NOT NULL CHECK ( octet_length(password) <> 8 ),
	about        						VARCHAR(150),
	avatar       						VARCHAR(200),
	user_type    						"level" 										NOT NULL DEFAULT 'user',
	"key"		 								VARCHAR(50)               	NOT NULL,
	active 									BOOLEAN 										NOT NULL  DEFAULT TRUE,
	attempts 								INTEGER 										NOT NULL  DEFAULT 0,
	last_failure 						TIMESTAMP,
	otp											BOOLEAN,
	created_by 							UUID,
	created_at   						TIMESTAMP WITH TIME ZONE    NOT NULL 	DEFAULT NOW(),
	last_login   						TIMESTAMP WITH TIME ZONE
);

CREATE INDEX users_email_idx ON public.users (email);
