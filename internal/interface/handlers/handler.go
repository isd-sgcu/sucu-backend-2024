package handlers

type Handler interface {
	Middleware() *MiddlewareHandler
	Auth() *AuthHandler
	User() *UserHandler
}
