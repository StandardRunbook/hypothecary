package shared

import (
	"context"

	plugininterface "github.com/StandardRunbook/plugin-interface/hypothesis-interface/github.com/StandardRunbook/hypothesis"
)

type GRPCServer struct {
	plugininterface.UnimplementedHypothesisServer
	Impl IPlugin
	cfg  map[string]string
}

func (g *GRPCServer) Init(_ context.Context, config *plugininterface.Config) (*plugininterface.InitResponse, error) {
	g.cfg = config.GetParameters()
	err := g.Impl.Init(g.cfg)
	if err != nil {
		return nil, err
	}
	return &plugininterface.InitResponse{
		ErrorMessage: ApplicationResponseSuccess,
	}, nil
}

func (g *GRPCServer) Name(_ context.Context, _ *plugininterface.Empty) (*plugininterface.NameResponse, error) {
	name, err := g.Impl.Name()
	if err != nil {
		return nil, err
	}
	return &plugininterface.NameResponse{
		Name: name,
	}, nil
}

func (g *GRPCServer) Version(_ context.Context, _ *plugininterface.Empty) (*plugininterface.VersionResponse, error) {
	version, err := g.Impl.Version()
	if err != nil {
		return nil, err
	}
	return &plugininterface.VersionResponse{
		Version: version,
	}, nil
}

func (g *GRPCServer) Run(_ context.Context, _ *plugininterface.Empty) (*plugininterface.RunResponse, error) {
	err := g.Impl.Run()
	if err != nil {
		return nil, err
	}
	return &plugininterface.RunResponse{
		ErrorMessage: ApplicationResponseSuccess,
	}, nil
}

func (g *GRPCServer) ParseOutput(_ context.Context, _ *plugininterface.Empty) (*plugininterface.ParseOutputResponse, error) {
	output, err := g.Impl.ParseOutput()
	if err != nil {
		return nil, err
	}
	return &plugininterface.ParseOutputResponse{
		Output: output,
	}, nil
}
