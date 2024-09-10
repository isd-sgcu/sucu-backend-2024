package repositories

type Repository interface {
	User() UserRepository
	Attachment() AttachmentRepository
	Document() DocumentRepository
}
