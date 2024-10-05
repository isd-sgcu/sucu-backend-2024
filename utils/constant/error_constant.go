package constant

var (
	// user error
	ErrUserAlreadyExists  = "user already exists"
	ErrUserNotFound       = "user not found"
	ErrHashPasswordFailed = "failed to hash password"
	ErrInsertUserFailed   = "failed to insert user"
	ErrFindUserByID       = "failed to find user by ID"
	ErrInvalidRole        = "invalid role"
	ErrUpdateUserByID     = "failed to update user by ID"

	// doc error
	ErrInvalidDocType       = "invalid document type"
	ErrInvalidOrg           = "invalid organization"
	ErrInsertDocumentFailed = "failed to insert document"
	ErrGetDocumentFailed    = "failed to get document"
	ErrInvalidTimeFormat    = "invalid time format"

	// pagination error
	ErrInvalidPageSize = "invalid page size"
)
