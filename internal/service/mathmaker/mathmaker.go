package mathmaker

import "go.uber.org/zap"

type Mathmaker struct {
	baseURL string
	logger  *zap.SugaredLogger
}

func New(baseURL string, logger *zap.SugaredLogger) *Mathmaker {
	return &Mathmaker{baseURL: baseURL, logger: logger}
}
