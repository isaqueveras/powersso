-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id			     						UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
	first_name   						VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
	last_name    						VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
	email        						VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
	password     						VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
	roles        						VARCHAR[]                 	NOT NULL DEFAULT '{}',
	about        						VARCHAR(500)                         DEFAULT '',
	avatar       						VARCHAR(512),
	phone_number 						VARCHAR(20),
	address      						VARCHAR(250),
	city         						VARCHAR(30),
	country      						VARCHAR(30),
	gender       						VARCHAR(20)                 					DEFAULT '',
	postcode     						INTEGER,
	token_key		 						VARCHAR(50)               					  DEFAULT '',
	registered_by 					UUID,
	is_active 							BOOLEAN 										NOT NULL  DEFAULT TRUE,
	number_failed_attempts 	INTEGER 										NOT NULL  DEFAULT 0,
	last_failure_date 			TIMESTAMP,
	birthday     						DATE                                 	DEFAULT NULL,
	created_at   						TIMESTAMP WITH TIME ZONE    NOT NULL 	DEFAULT NOW(),
	updated_at   						TIMESTAMP WITH TIME ZONE             	DEFAULT CURRENT_TIMESTAMP,
	login_date   						TIMESTAMP(0) WITH TIME ZONE NOT NULL 	DEFAULT CURRENT_TIMESTAMP
);
