package middles

import (
	"github.com/gofiber/fiber/v2"
)

// FiberRecoveryConfig defines the config for middleware.
type FiberRecoveryConfig struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// EnableStackTrace enables handling stack trace
	//
	// Optional. Default: false
	EnableStackTrace bool

	// StackTraceHandler defines a function to handle stack trace
	//
	// Optional. Default: defaultStackTraceHandler
	StackTraceHandler func(e interface{})
}

var defaultStackTraceBufLen = 1024

// DefaultFiberRecoveryConfig is the default config
var DefaultFiberRecoveryConfig = FiberRecoveryConfig{
	Next:              nil,
	EnableStackTrace:  false,
	StackTraceHandler: defaultStackTraceHandler,
}

// Helper function to set default values
func defaultFiberRecoveryConfig(config ...FiberRecoveryConfig) FiberRecoveryConfig {
	// Return default config if nothing provided
	if len(config) < 1 {
		return DefaultFiberRecoveryConfig
	}

	// Override default config
	cfg := config[0]

	if cfg.EnableStackTrace && cfg.StackTraceHandler == nil {
		cfg.StackTraceHandler = defaultStackTraceHandler
	}

	return cfg
}
