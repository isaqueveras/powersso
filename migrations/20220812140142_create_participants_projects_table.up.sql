-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE project_participants (
	"user_id"	      UUID NOT NULL REFERENCES users (id),
  "start_date"    TIMESTAMP WITH TIME ZONE NOT NULL,
  departure_date  TIMESTAMP WITH TIME ZONE,
	created_at	    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at	    TIMESTAMP WITH TIME ZONE
);
