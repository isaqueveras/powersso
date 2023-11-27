-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE projects (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" VARCHAR(30) NOT NULL,
	"desc" VARCHAR(50),
	"url" VARCHAR(200) NOT NULL,
	"created_by" UUID NOT NULL REFERENCES users (id),
	"updated_at" TIMESTAMP WITH TIME ZONE,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
