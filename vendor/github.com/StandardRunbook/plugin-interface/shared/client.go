package shared

import (
	"context"
	"errors"
	"fmt"
	"strings"

	plugininterface "github.com/StandardRunbook/plugin-interface/hypothesis-interface/github.com/StandardRunbook/hypothesis"
)

const (
	ApplicationResponseSuccess string = "no_error"
)

type GRPCClient struct {
	client plugininterface.HypothesisClient
}

func (g *GRPCClient) Init(m map[string]string) error {
	init, err := g.client.Init(context.Background(), &plugininterface.Config{
		Parameters: m,
	})
	if err != nil {
		return err
	}

	if strings.EqualFold(init.GetErrorMessage(), ApplicationResponseSuccess) {
		return nil
	}
	return errors.New(init.GetErrorMessage())
}

func (g *GRPCClient) Name() (string, error) {
	name, err := g.client.Name(context.Background(), nil)
	if err != nil {
		return "", err
	}
	if strings.EqualFold(name.GetName(), "") {
		return "", fmt.Errorf("name is empty")
	}
	return name.GetName(), nil
}

func (g *GRPCClient) Version() (string, error) {
	version, err := g.client.Version(context.Background(), nil)
	if err != nil {
		return "", err
	}
	if strings.EqualFold(version.GetVersion(), "") {
		return "", fmt.Errorf("version is empty")
	}
	return version.GetVersion(), nil
}

func (g *GRPCClient) Run() error {
	run, err := g.client.Run(context.Background(), nil)
	if err != nil {
		return err
	}

	if strings.EqualFold(run.GetErrorMessage(), ApplicationResponseSuccess) {
		return nil
	}
	return errors.New(run.GetErrorMessage())
}

func (g *GRPCClient) ParseOutput() (string, error) {
	output, err := g.client.ParseOutput(context.Background(), nil)
	if err != nil {
		return "", err
	}

	return output.GetOutput(), nil
}
