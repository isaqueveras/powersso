-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "level" AS ENUM ('user', 'admin', 'integration');

CREATE TABLE public."user" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"first_name" VARCHAR(20) NOT NULL CHECK ( first_name <> '' ),
	"last_name" VARCHAR(32) NOT NULL CHECK ( last_name <> '' ),
	"email" VARCHAR(64) UNIQUE NOT NULL CHECK ( email <> '' ),
	"password" VARCHAR(150) NOT NULL CHECK ( octet_length(password) > 8 ),
	"level" "level" NOT NULL DEFAULT 'user',
	"flag" INTEGER NOT NULL DEFAULT 0,
	"key" VARCHAR(50) NOT NULL,
	"active" BOOLEAN NOT NULL DEFAULT TRUE,
	"attempts" INTEGER NOT NULL DEFAULT 0,
	"last_failure" TIMESTAMP,
	"otp" VARCHAR,
	"created_by" UUID REFERENCES public."user" (id),
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	"last_login" TIMESTAMP WITH TIME ZONE
);

CREATE INDEX user_email_idx ON public.user (email);

CREATE TABLE public."session" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL REFERENCES public."user" (id),
	"ip" VARCHAR NOT NULL,
	"user_agent" VARCHAR NOT NULL,
	"expires_at" TIMESTAMP WITH TIME ZONE NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE public."organization" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" VARCHAR(30) NOT NULL,
	"desc" VARCHAR(50),
	"url" VARCHAR(200) NOT NULL,
	"created_by" UUID NOT NULL REFERENCES public."user" (id),
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE public."box" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" VARCHAR(30) NOT NULL,
	"desc" VARCHAR(50),
	"organization_id" UUID NOT NULL REFERENCES public."organization" (id),
	"permissions_id" UUID[] DEFAULT NULL,
	"users" UUID[] DEFAULT NULL,
	"created_by" UUID NOT NULL REFERENCES public."user" (id),
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE public."permission" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" VARCHAR(50) NOT NULL,
	"credential" VARCHAR(150) UNIQUE NOT NULL,
	"created_by" UUID NOT NULL REFERENCES public."user" (id),
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
