package main

import (
	"context"
	"os"
	"os/exec"

	"github.com/StandardRunbook/hypothecary/app/config"
	"github.com/StandardRunbook/plugin-interface/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog"
)

func run(binaryName string, cfg map[string]string) (string, error) {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", binaryName),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return "", err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("plugin_grpc")
	if err != nil {
		return "", err
	}

	hypothesis := raw.(shared.IPlugin)
	err = hypothesis.Init(cfg)
	if err != nil {
		return "", err
	}

	err = hypothesis.Run()
	if err != nil {
		return "", err
	}

	return hypothesis.ParseOutput()
}

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

	for _, i := range cfg.Plugins {

	}
	// once the config file is loaded, create the hypothesis dag

}
