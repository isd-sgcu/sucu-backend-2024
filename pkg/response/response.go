package response

import (
	"github.com/gofiber/fiber/v2"
)

const (
	SUCCESS string = "success"
	ERROR   string = "error"
)

type Response interface {
	SendResponse(c *fiber.Ctx, statusCode int) error
}

// Factory Method
func NewResponseFactory(responseType string, data interface{}) Response {
	switch responseType {
	case SUCCESS:
		return newSuccessResponse(data)
	case ERROR:
		return newErrorResponse(data)
	default:
		panic("Invalid response type")
	}
}
