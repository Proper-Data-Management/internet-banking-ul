package middles

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/mak-alex/al_hilal_core/helpers/apiErrors"
)

func defaultStackTraceHandler(e interface{}) {
	buf := make([]byte, defaultStackTraceBufLen)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}

// NewFiberRecovery creates a new middleware handler
func NewFiberRecovery(config ...FiberRecoveryConfig) fiber.Handler {
	// Set default config
	cfg := defaultFiberRecoveryConfig(config...)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				if cfg.EnableStackTrace {
					cfg.StackTraceHandler(r)
				}

				serverError := apiErrors.ThrowError(apiErrors.ServerError)
				_ = c.JSON(fiber.Map{
					"message": serverError.Message,
					"id":      serverError.Id,
					"status":  serverError.Status,
					"detail":  serverError.Detail,
				})
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
