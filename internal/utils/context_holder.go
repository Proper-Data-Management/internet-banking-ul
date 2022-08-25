package utils

import (
	"context"

	"sync"

	"github.com/internet-banking-ul/modules/logger"
)

// ContextHolderKey key for map in context
const ContextHolderKey = "ContextHolder"

// contextGetAttribute - get value from map stored in context
func contextGetAttribute(ctx context.Context, attribute string) interface{} {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		if value, ok := contextHolder.Load(attribute); ok {
			return value
		}
	}
	return nil
}

// SetAttribute - set value to map stored in context
func SetAttribute(ctx context.Context, attribute string, value interface{}) {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		contextHolder.Store(attribute, value)
	} else {
		logger.WorkLogger.Error("SetAttribute doesn't find context holder")
	}
}

// contextGetStringAttribute -
func contextGetStringAttribute(ctx context.Context, attribute string) (string, bool) {
	value := contextGetAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(string); ok {
			return result, true
		}
	}
	return "", false
}

// contextGetInt32Attribute -
func contextGetInt32Attribute(ctx context.Context, attribute string) (int32, bool) {
	value := contextGetAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(int32); ok {
			result := int32(result)
			return result, true
		}
	}
	return 0, false
}

// contextGetBoolAttribute -
func contextGetBoolAttribute(ctx context.Context, attribute string) (bool, bool) {
	value := contextGetAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(bool); ok {
			return result, true
		}
	}
	return false, false
}
