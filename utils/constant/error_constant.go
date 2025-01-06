package constant

var (
	// user error
	ErrUserAlreadyExists  = "user already exists"
	ErrUserNotFound       = "user not found"
	ErrRoleNotFound       = "role not found"
	ErrHashPasswordFailed = "failed to hash password"
	ErrInsertUserFailed   = "failed to insert user"
	ErrFindUserByID       = "failed to find user by ID"
	ErrInvalidRole        = "invalid role"
	ErrUpdateUserByID     = "failed to update user by ID"
	ErrDeleteUserByID     = "failed to delete user"
	ErrInvalidValue       = "invalid value"
	ErrInvalidQuery       = "invalid query"

	// doc error
	ErrInvalidDocType       = "invalid document type"
	ErrInvalidOrg           = "invalid organization"
	ErrInvalidTimeFormat    = "invalid time format"
	ErrDocumentNotFound     = "document not found"
	ErrFindDocumentByID     = "failed to find document by ID"
	ErrGetDocumentFailed    = "failed to get document"
	ErrInsertDocumentFailed = "failed to insert document"
	ErrUpdateDocumentFailed = "failed to update document"
	ErrDeleteDocumentFailed = "failed to delete document"

	// attachment error
	ErrAttachmentNotFound     = "attachment not found"
	ErrDeleteAttachmentFailed = "failed to delete attachment"
	ErrFindAttachmentByID     = "failed to find attachment by ID"
	// pagination error
	ErrInvalidPageSize = "invalid page size"
)
