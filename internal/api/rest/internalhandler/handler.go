package internalhandler

import (
	"go.uber.org/zap"
)

type App interface {
}

type Handler struct {
	app    App
	logger *zap.SugaredLogger
}

func New(app App, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		app:    app,
		logger: logger,
	}
}
