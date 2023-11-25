-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE projects (
	"id"			        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"created_by"	    UUID NOT NULL REFERENCES users (id),
	"name" 			      VARCHAR(20) NOT NULL,
	"description"     VARCHAR(50),
	"url"	            VARCHAR(200) NOT NULL,
  "slug"            VARCHAR(20) NOT NULL,
	"created_at"	    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	"updated_at"	    TIMESTAMP WITH TIME ZONE
);
