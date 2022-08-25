package logger

import (
	"context"
	"database/sql"
	"runtime/debug"

	"go.uber.org/zap"
)

func SqlWithCtxLevel(ctx context.Context, err error) func(msg string, fields ...zap.Field) {
	l := SqlLoggerWithContext(ctx).With(zap.Error(err))

	switch err {
	case nil:
		return l.Error
	case sql.ErrNoRows:
		return l.Warn
	default:
		return l.Panic
	}
}

func WorkWithCtxLevel(ctx context.Context, err error) func(msg string, fields ...zap.Field) {
	l := WorkLoggerWithContext(ctx).With(zap.Error(err))

	switch err {
	case nil:
		return l.Error
	case sql.ErrNoRows:
		return l.Warn
	default:
		return l.Panic
	}
}

// recovered - recover
// msg - additional message
func LogPanic(recovered interface{}, msg ...string) {
	WorkLogger.Error("panic_stack_trace", zap.Strings("msg", msg), zap.Any("recovered", recovered), zap.ByteString("stack", debug.Stack()))
}
