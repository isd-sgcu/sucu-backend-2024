package constant

import "errors"

var (
	// user error
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrHashPasswordFailed = errors.New("failed to hash password")
	ErrInsertUserFailed   = errors.New("failed to insert user")
	ErrFindUserByID       = errors.New("failed to find user by ID")
	ErrInvalidRole        = errors.New("invalid role")

	//document error
	ErrDocumentNotFound = errors.New("document not found")
	ErrFindDocumentByID = errors.New("failed to find document by ID")
	ErrUpdateDocumentFailed = errors.New("failed to update document")
	ErrDeleteDocumentFailed = errors.New("failed to delete document")

)
