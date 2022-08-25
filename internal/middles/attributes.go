package middles

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/mak-alex/al_hilal_core/internal/utils"
)

func SetupLanguage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Get("Translate-Language")
		if len(lang) == 0 || lang == "" {
			lang = "RU"
		}
		if iHolder := c.Locals(utils.ContextHolderKey); iHolder != nil {
			if holder, ok := iHolder.(*sync.Map); ok {
				holder.Store("locale", lang)
			}
		}
		return c.Next()
	}
}

func SetupRequestInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if iHolder := c.Locals(utils.ContextHolderKey); iHolder != nil {
			if holder, ok := iHolder.(*sync.Map); ok {
				//holder[utils.AttributeIsMobile] = isMobile(c.Fasthttp.Request)
				holder.Store(utils.AttributeUserAgent, c.Get("User-Agent"))
				holder.Store(utils.AttributeAppID, c.Get("X-DigitalBank-application-id"))
				holder.Store(utils.AttributeAppName, c.Get("X-DigitalBank-app-name"))
				holder.Store(utils.AttributeDeviceID, c.Get("X-DigitalBank-device-id"))
				holder.Store(utils.AttributeXRealIP, c.Get("X-Real-Ip"))
				holder.Store(utils.AttributeXForwardedFor, c.Get("X-Forwarded-For"))
				holder.Store(utils.AttributeXOriginalForwardedFor, c.Get("X-Original-Forwarded-For"))
				_, os, appVersion := utils.ParseUserAgent(c.Get("User-Agent"))
				holder.Store(utils.AttributeAppVersion, appVersion)
				holder.Store(utils.AttributeOperationSystem, os)
			}
		}
		return c.Next()
	}
}
