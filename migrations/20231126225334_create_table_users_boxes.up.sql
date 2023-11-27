-- Copyright (c) 2023 Isaque Veras
-- Use of this source code is governed by MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE users_boxes (
	id UUID  DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES users (id),
	box_id UUID NOT NULL REFERENCES boxes (id),
	created_by UUID NOT NULL REFERENCES users (id),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
);
