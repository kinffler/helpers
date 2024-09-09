package helpersResponse

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// DefaultResponse represents a default HTTP response
// DefaultResponse godoc
// @Description Default response structure
type DefaultResponse struct {
	Message string `json:"message,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

func SuccessCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(data)
}

func BadRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusBadRequest).JSON(DefaultResponse{
		Message: message,
	})
}

func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(DefaultResponse{
		Message: message,
	})
}

func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusForbidden).JSON(DefaultResponse{
		Message: message,
	})
}

func NotFoundResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusNotFound).JSON(DefaultResponse{
		Message: message,
	})
}

func StatusConflict(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusConflict).JSON(DefaultResponse{
		Message: message,
	})
}

func UnprocessableResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(DefaultResponse{
		Message: message,
	})
}

func TooManyRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusTooManyRequests).JSON(DefaultResponse{
		Message: message,
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusInternalServerError).JSON(DefaultResponse{
		Message: message,
	})
}
