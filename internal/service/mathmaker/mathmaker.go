package mathmaker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Mathmaker struct {
	baseURL string
	client  http.Client
	logger  *zap.SugaredLogger
}

func New(baseURL string, reqTimeout time.Duration, logger *zap.SugaredLogger) *Mathmaker {
	client := http.Client{
		Timeout: reqTimeout,
	}

	return &Mathmaker{baseURL: baseURL, client: client, logger: logger}
}

func (m *Mathmaker) Get(ctx context.Context, url string) ([]byte, error) {
	const op = "mathmaker.mathmaker.Get"

	m.logger.Debugf("url", url)

	resp, err := m.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	m.logger.With("url", url).Debugf("resp status code", resp.StatusCode)

	if err := checkStatusCode(resp.StatusCode); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return body, nil
}

func checkStatusCode(code int) error {
	const op = "mathmaker.mathmaker.checkStatusCode"

	switch code {
	case http.StatusNotFound:
		return fmt.Errorf("%s: %w", op, errors.New("not found"))
	}

	return nil
}
