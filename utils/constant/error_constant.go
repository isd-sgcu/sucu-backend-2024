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
	ErrInvalidAttachmentType = "invalid attachment type"
	ErrGetAttachmentFailed    = "failed to get attachment"


	// pagination error
	ErrInvalidPageSize = "invalid page size"
)
