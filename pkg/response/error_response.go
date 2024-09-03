package response

import "github.com/gofiber/fiber/v2"

// Error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func newErrorResponse(data interface{}) *ErrorResponse {
	message, ok := data.(string)
	if !ok {
		message = "an error occured"
	}

	return &ErrorResponse{
		Success: false,
		Message: message,
	}
}

func (r *ErrorResponse) SendResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(r)
}
