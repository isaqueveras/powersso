-- Copyright (c) 2022 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE project_participants (
	project_id	  	UUID NOT NULL REFERENCES projects (id),
	user_id	      	UUID NOT NULL REFERENCES users (id),
  "start_date"    DATE NOT NULL,
  departure_date  DATE,
	created_at	    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at	    TIMESTAMP WITH TIME ZONE
);
