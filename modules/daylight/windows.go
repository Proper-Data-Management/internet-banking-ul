//go:build windows
// +build windows

package daylight

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/mak-alex/al_hilal_core/modules/logger"
	"go.uber.org/zap"
)

// KillPid kills the process with the specified pid
func KillPid(ctx context.Context, pid string) error {
	l := logger.WorkLoggerWithContext(ctx)
	rez, err := exec.Command("tasklist", "/fi", "PID eq "+pid).Output()
	if err != nil {
		l.With(zap.String("type", "CommandExecutionError"), zap.Error(err), zap.String("cmd", "tasklist /fi PID eq"+pid)).Error("Error executing command")
		return err
	}
	if string(rez) == "" {
		return fmt.Errorf("null")
	}

	l.With(zap.String("cmd", "tasklist /fi PID eq "+pid)).Debug("command execution result")
	if ok, _ := regexp.MatchString(`(?i)PID`, string(rez)); !ok {
		return fmt.Errorf("null")
	}
	return nil
}
