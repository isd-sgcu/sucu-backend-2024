package usecases

type Usecase interface {
	Middleware() MiddlewareUsecase
	Auth() AuthUsecase
	User() UserUsecase
	Attachment() AttachmentUsecase
	Document() DocumentUsecase
}
