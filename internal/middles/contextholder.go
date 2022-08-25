package middles

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/internet-banking-ul/internal/utils"
)

func SetupContextHolder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(utils.ContextHolderKey, &sync.Map{})
		return c.Next()
	}
}
