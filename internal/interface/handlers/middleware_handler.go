package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type MiddlewareHandler struct {
	middlewareUsecase usecases.MiddlewareUsecase
}

func NewMiddlewareHandler(middlewareUsecase usecases.MiddlewareUsecase) *MiddlewareHandler {
	return &MiddlewareHandler{
		middlewareUsecase: middlewareUsecase,
	}
}

func (h *MiddlewareHandler) IsLogin(c *fiber.Ctx) error {
	// get header "Authorization"
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		resp := response.NewResponseFactory(response.ERROR, errors.New("Unauthorized").Error())
		return resp.SendResponse(c, fiber.StatusUnauthorized)
	}

	// validate Bearer token
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		resp := response.NewResponseFactory(response.ERROR, errors.New("Unauthorized").Error())
		return resp.SendResponse(c, fiber.StatusUnauthorized)
	}

	// get token
	token := authHeader[len(bearerPrefix):]

	// verify token
	userID, err := h.middlewareUsecase.VerifyToken(token)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, errors.New("Unauthorized").Error())
		return resp.SendResponse(c, fiber.StatusUnauthorized)
	}

	// get requested user data
	userDTO, err := h.middlewareUsecase.GetMe(*userID)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, errors.New("Unauthorized").Error())
		return resp.SendResponse(c, fiber.StatusUnauthorized)
	}

	// store userDTO in context
	c.Locals("user", userDTO)

	// move to next handlers
	return c.Next()
}
