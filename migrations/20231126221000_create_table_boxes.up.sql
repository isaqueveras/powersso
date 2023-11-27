-- Copyright (c) 2023 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE boxes (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" VARCHAR(30) NOT NULL,
	"desc" VARCHAR(50),
	"project_id" UUID NOT NULL REFERENCES projects (id),
	"flag" BIGINT NOT NULL DEFAULT 0,
	"created_by" UUID NOT NULL REFERENCES users (id),
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
);
