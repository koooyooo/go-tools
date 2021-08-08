package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestSugarLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "http://locallhost:oooo/",
		"attempt", 3,
		"backoff", time.Second,
	)
}

func TestLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("failed to fetch URL",
		zap.String("url", "http://locallhost:oooo/"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// BenchmarkSugarLogger-8   	 5147736	       239.4 ns/op
func BenchmarkSugarLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	for i := 0; i < b.N; i++ {
		sugar.Infow("failed to fetch URL",
			"url", "http://locallhost:oooo/",
			"attempt", 3,
			"backoff", time.Second,
		)
	}
}

// BenchmarkLogger-8   	 4074379	       310.9 ns/op
func BenchmarkLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL",
			zap.String("url", "http://locallhost:oooo/"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
