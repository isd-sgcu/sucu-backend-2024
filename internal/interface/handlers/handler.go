package handlers

type Handler interface {
	Middleware() *MiddlewareHandler
	Auth() *AuthHandler
	User() *UserHandler
	Attachment() *AttachmentHandler
	Document() *DocumentHandler
}
