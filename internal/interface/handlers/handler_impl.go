package handlers

import "github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"

type handler struct {
	MiddlewareHandler *MiddlewareHandler
	AuthHandler       *AuthHandler
	UserHandler       *UserHandler
	AttachmentHandler *AttachmentHandler
}

func NewHandler(usecases usecases.Usecase) Handler {
	return &handler{
		MiddlewareHandler: NewMiddlewareHandler(usecases.Middleware()),
		AuthHandler:       NewAuthHandler(usecases.Auth()),
		UserHandler:       NewUserHandler(usecases.User()),
		AttachmentHandler: NewAttachmentHandler(usecases.Attachment()),
	}
}

func (h *handler) Middleware() *MiddlewareHandler {
	return h.MiddlewareHandler
}

func (h *handler) Auth() *AuthHandler {
	return h.AuthHandler
}

func (h *handler) User() *UserHandler {
	return h.UserHandler
}

func (h *handler) Attachment() *AttachmentHandler {
	return h.AttachmentHandler
}
