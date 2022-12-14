package middles

import (
	"time"

	logger2 "github.com/internet-banking-ul/modules/logger"

	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
)

//Logging colors, unused until zap implements colored logging -> https://github.com/uber-go/zap/issues/489
var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

//setupLogging setups the logger to use zap
func setupLogging(duration time.Duration, zap *zap.Logger) {
	go func() {
		for range time.Tick(duration) {
			zap.Sync()
		}
	}()
}

//Logger returns a fiber handler func for all logging
func Logger(duration time.Duration, logger *zap.Logger) fiber.Handler {
	setupLogging(duration, logger)
	return func(c *fiber.Ctx) (err error) {
		t := time.Now()
		defer func() {
			if err := recover(); err != nil {
				logger2.LogPanic(err)
				panic(err)
			}
		}()
		err = c.Next()
		var errStr string
		if err != nil {
			errStr = err.Error()
		}

		latency := time.Since(t)
		clientIP := c.IP()
		method := c.Method()
		statusCode := c.Response().StatusCode()
		path := c.Path()

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				logger.Warn("[Fiber]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("error", errStr),
				)
			}
		case statusCode >= 500:
			{
				logger.Error("[Fiber]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("error", errStr),
				)
			}
		default:
			if err != nil {

			}
			logger.Info("[Fiber]",
				zap.Int("statusCode", statusCode),
				zap.String("latency", latency.String()),
				zap.String("clientIP", clientIP),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("error", errStr),
			)
		}
		return err
	}
}
