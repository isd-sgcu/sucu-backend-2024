package repositories

type Repository interface {
	User() UserRepository
	Attachment() AttachmentRepository
}
