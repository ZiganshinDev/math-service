package mathmaker

import (
	"context"
	"fmt"
	"io"
	"mathbot/internal/svcerr"
	"net/http"
	"strings"
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

	m.logger.Debugf("requesting URL: %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	header := http.Header{
		"accept":          []string{"application/json"},
		"accept-encoding": []string{"gzip", "deflate", "br", "zstd"},
	}

	req.Header = header

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, svcerr.New(svcerr.ErrInternal, err))
	}
	defer resp.Body.Close()

	m.logger.With("url", url).Debugf("response status: %d, headers: %v", resp.StatusCode, resp.Header)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		m.logger.Debugf("response body: %s", body)
		return nil, nil
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		m.logger.Warnf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, svcerr.New(svcerr.ErrInternal, err))
	}

	m.logger.Debugf("response body length: %d bytes", len(body))
	return body, nil
}
