package apiErrors

import "github.com/gofiber/fiber/v2"

func Send(c *fiber.Ctx, errorID string) error {
	apiErr := ThrowError(errorID)
	return SendErr(c, apiErr)
}

func SendErr(c *fiber.Ctx, err *apiError) error {
	return c.Status(err.Status).JSON(fiber.Map{
		"error": err.Message,
	})
}
