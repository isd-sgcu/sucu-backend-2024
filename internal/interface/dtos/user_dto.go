package dtos

import "time"

type UserDTO struct {
	ID        string    `json:"id"`         // student id
	FirstName string    `json:"first_name"` // user's first name
	LastName  string    `json:"last_name"`  // user's last name
	Role      string    `json:"role"`       // role: sgcu-admin, sgcu-superadmin, sccu-admin, sccu-superadmin
	CreatedAt time.Time `json:"created_at"` // user's account creation time
	UpdatedAt time.Time `json:"updated_at"` // user's last update time
}

type CreateUserDTO struct {
	ID        string `json:"id" validate:"required"`         // student id
	FirstName string `json:"first_name" validate:"required"` // user's first name
	LastName  string `json:"last_name" validate:"required"`  // user's last name
	Password  string `json:"password" validate:"required"`   // user's password
	Role      string `json:"role"`                           // role: sgcu-admin, sgcu-superadmin, sccu-admin, sccu-superadmin
}

type UpdateUserDTO struct {
	FirstName string `json:"first_name"` // user's first name
	LastName  string `json:"last_name"`  // user's last name
	Password  string `json:"password"`   // user's password
}

type GetAllUsersDTO struct {
	Page  int `json:"page"`  // current page
	Limit int `json:"limit"` // page size
}
