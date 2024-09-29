package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"go.uber.org/zap"
)

type usecase struct {
	MiddlewareUsecase MiddlewareUsecase
	AuthUsecase       AuthUsecase
	UserUsecase       UserUsecase
	AttachmentUsecase AttachmentUsecase
	DocumentUsecase   DocumentUsecase
}

func NewUsecase(repo repositories.Repository, cfg config.Config, logger *zap.Logger) Usecase {
	return &usecase{
		MiddlewareUsecase: NewMiddlewareUsecase(cfg, logger.Named("MiddlewareSvc"), repo.User()),
		AuthUsecase:       NewAuthUsecase(cfg, logger.Named("AuthSvc"), repo.User()),
		UserUsecase:       NewUserUsecase(cfg, logger.Named("UserSvc"), repo.User()),
		AttachmentUsecase: NewAttachmentUsecase(cfg, logger.Named("AttachmentSvc"), repo.Attachment()),
		DocumentUsecase:   NewDocumentUsecase(cfg, logger.Named("DocumentSvc"), repo.Document(), repo.User()),
	}
}

func (u *usecase) Middleware() MiddlewareUsecase {
	return u.MiddlewareUsecase
}

func (u *usecase) Auth() AuthUsecase {
	return u.AuthUsecase
}

func (u *usecase) User() UserUsecase {
	return u.UserUsecase
}

func (u *usecase) Attachment() AttachmentUsecase {
	return u.AttachmentUsecase
}

func (u *usecase) Document() DocumentUsecase {
	return u.DocumentUsecase
}
