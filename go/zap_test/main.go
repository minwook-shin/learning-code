package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Debugw("Debugw",
		"attempt", 3,
		"time", time.Second,
	)
	sugar.Debugf("Debugf")

	sugar.Infow("Infow",
		"attempt", 3,
		"time", time.Second,
	)
	sugar.Infof("Infof")

	sugar.Warnw("Warnw",
		"attempt", 3,
		"time", time.Second,
	)
	sugar.Warnf("Warnf")

	sugar.Errorw("Errorw",
		"attempt", 3,
		"time", time.Second,
	)
	sugar.Errorf("Errorf")

	// sugar.Fatalw("Fatalw",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Fatalf("Fatalf")

	// sugar.Panicw("Panicw",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Panicf("Panicf")

	logger.Debug("Debug",
		zap.String("strKey", "String"),
		zap.Int("intKey", 1),
		zap.Duration("Duration", time.Second),
	)

	logger.Info("Info",
		zap.String("strKey", "String"),
		zap.Duration("Duration", time.Second),
	)

	logger.Warn("Warn",
		zap.String("strKey", "String"),
		zap.Duration("Duration", time.Second),
	)

	logger.Error("Error",
		zap.String("strKey", "String"),
		zap.Duration("Duration", time.Second),
	)

	logger.Fatal("Fatal",
		zap.String("strKey", "String"),
		zap.Duration("Duration", time.Second),
	)

	logger.Panic("Panic",
		zap.String("strKey", "String"),
		zap.Duration("Duration", time.Second),
	)

}
