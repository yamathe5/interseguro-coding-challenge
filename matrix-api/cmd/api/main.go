package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	api "interseguro/matrix-api/internal/adapters/http"
	"interseguro/matrix-api/internal/adapters/statistics"
)

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	port := env("PORT", "8080")
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		log.Fatal("Invalid PORT")
	}
	if len(os.Args) == 2 && os.Args[1] == "healthcheck" {
		client := http.Client{Timeout: 2 * time.Second}
		res, err := client.Get(fmt.Sprintf("http://127.0.0.1:%s/health/live", port))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			log.Fatalf("HTTP %d", res.StatusCode)
		}
		return
	}
	timeout, err := time.ParseDuration(env("STATISTICS_TIMEOUT", "3s"))
	if err != nil {
		log.Fatal("Invalid STATISTICS_TIMEOUT")
	}
	client, err := statistics.New(env("STATISTICS_URL", "http://127.0.0.1:3000"), timeout)
	if err != nil {
		log.Fatal(err)
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("listening", "port", port)
	if err := api.New(client).Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
