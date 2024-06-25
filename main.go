package main

import (
	"context"
	"github.com/StandardRunbook/hypothecary/app/config"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	ctx := context.Background()
	logger := zerolog.Ctx(ctx)
	ctx = logger.WithContext(ctx)
	configFile, exists := os.LookupEnv("HYPOTHECARY_CONFIG_FILE")
	if !exists {
		configFile = "config.yaml"
	}
	// load the config file
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load config file")
	}

	// once the config file is loaded, create the hypothesis dag

}
