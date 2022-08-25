package daylight

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"log"

	"github.com/mak-alex/al_hilal_core/internal/config"
	"github.com/mak-alex/al_hilal_core/internal/server"
	"github.com/mak-alex/al_hilal_core/modules/logger"
	"github.com/mak-alex/al_hilal_core/tools"
	"go.uber.org/zap"
)

func Start(ctx context.Context) {
	var err error

	defer func() {
		r := recover()
		if r != nil {
			var err error
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("unknown error")
			}
			// sendMeMail(err)
			log.Fatalln(err)
		}
	}()

	cfg := config.Load()
	if cfg == nil {
		os.Exit(1)
	}

	Exit := func(code int) {
		delPidFile(cfg)
		os.Exit(code)
	}

	if err := tools.MakeDirectory(cfg.DataDir); err != nil {
		log.Fatal("can't create temporary directory")
		Exit(1)
	}

	isDebug := strings.Contains(cfg.Environment, "dev")
	logger.Init(
		isDebug,
		&logger.Config{
			EnableConsole:     true,
			EnableFile:        len(cfg.LogFile) > 0,
			ConsoleJSONFormat: false,
			ConsoleLevel:      tools.Ternary(isDebug, "debug", "info"),
			FileJSONFormat:    true,
			FileLevel:         cfg.LogLevel,
			FileLocation:      cfg.LogFile,
			FileMaxSize:       2,
			FileMaxBackups:    2,
			FileMaxAge:        2,
			FileCompress:      true,
		},
	)

	l := logger.WorkLoggerWithContext(ctx)

	sqlDB, err := sql.Open("oci8", cfg.DSN())
	if err != nil {
		l.Error("Failed close DB connection", zap.Error(err))
		Exit(1)
	}
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			l.Error("Failed close DB connection", zap.Error(err))
			Exit(1)
		}
	}()

	err = sqlDB.PingContext(ctx)
	if err != nil {
		l.Error("Failed close DB connection", zap.Error(err))
		Exit(1)
	}

	f := tools.LockOrDie(ctx, cfg.LockFilePath)
	defer f.Unlock()

	if err := tools.MakeDirectory(cfg.TempDir); err != nil {
		l.
			With(
				zap.Error(err),
				zap.String("type", "IOError"),
				zap.String("dir", cfg.TempDir),
			).
			Error("can't create temporary directory")
		Exit(1)
	}

	killOld(ctx, cfg)

	rand.Seed(time.Now().UTC().UnixNano())

	// save the current pid and version
	if err := savePid(ctx, cfg); err != nil {
		log.Fatalf("can't create pid: %s", err)
		Exit(1)
	}
	defer delPidFile(cfg)

	logger.WorkLoggerWithContext(ctx).Info("al_hilal_core started")

	srv := server.NewServer(sqlDB)
	if err := srv.Listen(cfg.ServerPort); err != nil {
		log.Panic(err)
	}

	go initGracefulShutDown(cfg)
}

func initGracefulShutDown(
	cfg *config.Config,
) {
	sigChan := make(chan os.Signal, 1)
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Exit(2)
	}()
}

func savePid(ctx context.Context, cfg *config.Config) error {
	if cfg == nil {
		return fmt.Errorf("`cfg' can't be empty")
	}
	pid := os.Getpid()
	pidAndVer, err := json.Marshal(map[string]string{
		"pid":     tools.ToStr(pid),
		"version": "0.0.1",
	})
	if err != nil {
		logger.
			WorkLoggerWithContext(ctx).
			With(
				zap.Any("pid", pid),
				zap.Error(err),
				zap.String("type", "JSONMarshallError"),
			).
			Error("marshalling pid to json")
		return err
	}

	return ioutil.WriteFile(cfg.GetPidPath(), pidAndVer, 0644)
}

func delPidFile(cfg *config.Config) {
	if cfg == nil {
		return
	}
	os.Remove(cfg.GetPidPath())
}

func killOld(ctx context.Context, cfg *config.Config) {
	if cfg == nil {
		return
	}

	l := logger.WorkLoggerWithContext(ctx)
	pidPath := cfg.GetPidPath()
	if _, err := os.Stat(pidPath); err == nil {
		dat, err := ioutil.ReadFile(pidPath)
		if err != nil {
			l.
				With(
					zap.String("path", pidPath),
					zap.Error(err),
					zap.String("type", "IOError"),
				).
				Error("reading pid file")
		}

		var pidMap map[string]string
		err = json.Unmarshal(dat, &pidMap)
		if err != nil {
			l.
				With(
					zap.ByteString("data", dat),
					zap.Error(err),
					zap.String("type", "JSONUnmarshallError"),
				).
				Error("unmarshalling pid map")
		}

		if pidMap["pid"] == "" {
			return
		}
		KillPid(ctx, pidMap["pid"])
		if fmt.Sprintf("%s", err) != "null" {
			// give 15 sec to end the previous process
			for i := 0; i < 15; i++ {
				if _, err := os.Stat(cfg.GetPidPath()); err == nil {
					time.Sleep(time.Second)
				} else {
					break
				}
			}
		}
	}
}
