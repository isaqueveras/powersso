-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE activate_account_tokens (
	id			     UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id			 UUID NOT NULL REFERENCES users (id),
	used 				 BOOLEAN NOT NULL DEFAULT FALSE,
	expires_at	 TIMESTAMP WITH TIME ZONE NOT NULL,
	created_at	 TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at	 TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
