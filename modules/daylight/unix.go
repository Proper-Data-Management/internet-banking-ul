//go:build (linux || freebsd || darwin) && (386 || amd64)
// +build linux freebsd darwin
// +build 386 amd64

package daylight

import (
	"context"
	"errors"
	"syscall"

	"github.com/mak-alex/al_hilal_core/modules/logger"
	"github.com/mak-alex/al_hilal_core/tools"
	"go.uber.org/zap"
)

// KillPid is killing process by PID
func KillPid(ctx context.Context, pid string) error {
	l := logger.
		WorkLoggerWithContext(ctx).
		With(zap.String("pid", pid))
	err := syscall.Kill(tools.ToInt(pid), syscall.SIGTERM)
	if err != nil {
		// http://www.gnu.org/software/libc/manual/html_node/Error-Codes.html
		if errors.Is(err, syscall.ESRCH) {
			l.
				With(zap.Any("detail", syscall.ESRCH)).
				Debug("process killed")
		} else {
			l.
				With(zap.Any("signal", syscall.SIGTERM)).
				Error("Error killing process with pid", zap.Error(err))
		}

		return err
	}
	return nil
}
