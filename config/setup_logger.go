package config

import "LoanTrackerApi/internal/logging"

var Logger logging.Logger

func SetupLogger() {
	Logger = logging.NewInMemoryLogManager()
}
