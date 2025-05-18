package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mathbot/internal/api/rest/server"
	"mathbot/internal/config"
	"mathbot/internal/service/app"
	"mathbot/internal/service/mathmaker"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

const (
	prod = "prod"
)

func main() {
	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	l, err := initLogger(cfg.Env)
	logger := l.Sugar()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Print(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("logger creating error %v", err)
	}

	logger.Info(slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	// mathmaker
	mathmaker := mathmaker.New(cfg.Mathmaker.BaseURL, cfg.Mathmaker.Timeout, logger)
	_, err = mathmaker.Problems(ctx)
	if err != nil {
		fmt.Println("error", err)
	}
	res, err := mathmaker.Problem(ctx, uuid.MustParse("0a76cc73-ea95-4c9b-8b15-d1ee21de8aaa"))
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("res", res)

	// go
	g, gCtx := errgroup.WithContext(ctx)

	// app
	mathbot := app.New(mathmaker)

	// rest server
	rest := server.New(cfg.HTTPServer.Port, mathbot, logger)

	logger.Info("starting REST server")
	g.Go(func() error {
		err := rest.Start()
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	})
	logger.Info("REST server started")

	g.Go(func() error {
		<-gCtx.Done()
		logger.Info("stopping REST server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return rest.Shutdown(shutdownCtx)
	})

	if err = g.Wait(); err != nil {
		logger.Error("errors from errorGroup: %v", err.Error())
	}

	logger.Info("REST server stopped")
	logger.Info("main done")
}

func initLogger(env string) (*zap.Logger, error) {
	var zapConfig zap.Config

	switch env {
	case prod:
		zapConfig = zap.NewProductionConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	zapConfig.EncoderConfig.TimeKey = "time"
	l, err := zapConfig.Build()

	return l, err
}
