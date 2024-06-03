package core

import (
	"fmt"

	"go.uber.org/zap"
)

var lg *zap.Logger

func ConfigureLogger() error {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("failed configuring zap logger: %w", err)
	}
	lg = zapLogger
	return nil
}

func GetLogger() *zap.Logger {
	return lg
}
