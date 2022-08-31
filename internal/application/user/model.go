package user

import "time"

// User is the model for user
type User struct {
	ID          *string    `json:"id,omitempty"`
	Email       *string    `json:"email,omitempty"`
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	Roles       []string   `json:"roles,omitempty"`
	About       *string    `json:"about,omitempty"`
	Avatar      *string    `json:"avatar,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Address     *string    `json:"address,omitempty"`
	City        *string    `json:"city,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Postcode    *int       `json:"postcode,omitempty"`
	TokenKey    *string    `json:"-"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	LoginDate   *time.Time `json:"login_date,omitempty"`
}
