package response

import "github.com/gofiber/fiber/v2"

// Success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
}

func newSuccessResponse(data interface{}) *SuccessResponse {
	if data == nil {
		data = struct{}{}
	}

	return &SuccessResponse{
		Success: true,
		Result:  data,
	}
}

func (r *SuccessResponse) SendResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(r)
}
