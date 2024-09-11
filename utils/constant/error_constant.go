package constant

import "errors"

var (
	// user error
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrHashPasswordFailed = errors.New("failed to hash password")
	ErrInsertUserFailed   = errors.New("failed to insert user")
	ErrFindUserByID       = errors.New("failed to find user by ID")
)
