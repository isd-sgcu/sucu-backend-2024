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
	ErrInsertDocumentFailed = "failed to insert document"
	ErrDocumentNotFound = "document not found"
	ErrFindDocumentByID = "failed to find document by ID"
	ErrUpdateDocumentFailed = "failed to update document"
	ErrDeleteDocumentFailed = "failed to delete document"
)
