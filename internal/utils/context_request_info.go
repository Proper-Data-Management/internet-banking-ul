package utils

import (
	"context"
	"strings"
)

const (
	AttributeLocale                = "locale"
	AttributeUserAgent             = "user_agent"
	AttributeIsMobile              = "is_mobile"
	AttributeAppID                 = "app_id"
	AttributeAppName               = "app_name"
	AttributeAppVersion            = "app_version"
	AttributeDeviceID              = "device_id"
	AttributeOperationSystem       = "operation_system"
	AttributeTerminalIP            = "terminal_ip"
	AttributeXForwardedFor         = "x_forwarded_for"
	AttributeXOriginalForwardedFor = "x_original_forwarded_for"
	AttributeXRealIP               = "x_real_ip"
	AttributeCurrentCtnInd         = "current_ctn_ind"
	AttributeCurrentEmail          = "current_email"
	AttributeCurrentUsername       = "current_username"
	AttributeCurrentUserID         = "current_user_id"
	AttributeCurrentUser           = "current_user"
)

func ContextGetLocale(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeLocale)
}

func ContextGetCurrentCtnInd(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeCurrentCtnInd)
}

func ContextGetCurrentEmail(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeCurrentEmail)
}

func ContextGetCurrentUsername(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeCurrentUsername)
}

func ContextGetCurrentUserID(ctx context.Context) (int32, bool) {
	return contextGetInt32Attribute(ctx, AttributeCurrentUserID)
}

//func ContextGetCurrentUser(ctx context.Context) *models.Customers {
//	result := contextGetAttribute(ctx, AttributeCurrentUser)
//	if result != nil {
//		currentUser := result.(models.Customers)
//		return &currentUser
//	}
//	return nil
//}

func ContextGetUserAgent(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeUserAgent)
}

func ContextGetIsMobile(ctx context.Context) (bool, bool) {
	return contextGetBoolAttribute(ctx, AttributeIsMobile)
}

func ContextGetAppID(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeAppID)
}

func ContextGetAppName(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeAppName)
}

func ContextGetAppVersion(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeAppVersion)
}

func ContextGetDeviceID(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeDeviceID)
}

func ContextGetOperationSystem(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeOperationSystem)
}

func ContextGetTerminalIP(ctx context.Context) (string, bool) {
	return contextGetStringAttribute(ctx, AttributeTerminalIP)
}

func ContextGetLocalIP(ctx context.Context) (ip string) {
	ip = "127.0.0.1"
	if xOriginalForwardedFor, ok := contextGetStringAttribute(ctx, AttributeXOriginalForwardedFor); ok {
		ip = strings.TrimSpace(strings.Split(xOriginalForwardedFor, ",")[0])
		return
	}

	if xForwardedFor, ok := contextGetStringAttribute(ctx, AttributeXForwardedFor); ok {
		ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
		return
	}

	if xRealIp, ok := contextGetStringAttribute(ctx, AttributeXRealIP); ok {
		ip = strings.TrimSpace(xRealIp)
		return
	}

	return
}
