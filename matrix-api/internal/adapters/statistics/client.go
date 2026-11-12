package statistics

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"interseguro/matrix-api/internal/application"
)

type Client struct {
	base string
	http *http.Client
}

func New(base string, timeout time.Duration) (*Client, error) {
	parsed, err := url.Parse(base)
	if err != nil || parsed.Hostname() == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.User != nil ||
		parsed.RawQuery != "" || parsed.ForceQuery || parsed.Fragment != "" || (parsed.Path != "" && parsed.Path != "/") || timeout <= 0 || timeout > 5*time.Second {
		return nil, errors.New("Invalid statistics URL or timeout (must be >0 and <=5s)")
	}
	return &Client{strings.TrimRight(base, "/"), &http.Client{Timeout: timeout, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}, nil
}

func (c *Client) Ready(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/health/ready", nil)
	if err != nil {
		return err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("statistics health: HTTP %d", res.StatusCode)
	}
	return nil
}

func (c *Client) Analyze(ctx context.Context, matrices [][][]float64, id string) (application.Statistics, error) {
	body, err := json.Marshal(map[string][][][]float64{"matrices": matrices})
	if err != nil {
		return application.Statistics{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/internal/v1/statistics", bytes.NewReader(body))
	if err != nil {
		return application.Statistics{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", id)
	res, err := c.http.Do(req)
	if err != nil {
		return application.Statistics{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return application.Statistics{}, fmt.Errorf("statistics: HTTP %d", res.StatusCode)
	}
	const maxResponse = 16 << 10
	data, err := io.ReadAll(io.LimitReader(res.Body, maxResponse+1))
	if err != nil || len(data) > maxResponse {
		return application.Statistics{}, errors.New("invalid response body")
	}
	var wire struct {
		Max, Min, Average, Sum *float64
		AnyDiagonal            *bool
		DiagonalMatrices       []int
		MatrixCount            int
	}
	if json.Unmarshal(data, &wire) != nil || wire.Max == nil || wire.Min == nil || wire.Average == nil || wire.Sum == nil ||
		wire.AnyDiagonal == nil || wire.DiagonalMatrices == nil || wire.MatrixCount != len(matrices) ||
		*wire.AnyDiagonal != (len(wire.DiagonalMatrices) > 0) {
		return application.Statistics{}, errors.New("invalid statistics response")
	}
	previous := -1
	for _, index := range wire.DiagonalMatrices {
		if index <= previous || index >= len(matrices) {
			return application.Statistics{}, errors.New("invalid diagonal matrix index")
		}
		previous = index
	}
	return application.Statistics{Max: *wire.Max, Min: *wire.Min, Average: *wire.Average, Sum: *wire.Sum,
		AnyDiagonal: *wire.AnyDiagonal, DiagonalMatrices: wire.DiagonalMatrices, MatrixCount: wire.MatrixCount}, nil
}
