package shared

import (
	"context"
	_ "net/rpc"

	"google.golang.org/grpc"

	proto "github.com/StandardRunbook/plugin-interface/hypothesis-interface/github.com/StandardRunbook/hypothesis"
	"github.com/hashicorp/go-plugin"
)

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"plugin_grpc": &GRPCPlugin{},
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

type IPlugin interface {
	Init(map[string]string) error
	Name() (string, error)
	Version() (string, error)
	Run() error
	ParseOutput() (string, error)
}

type GRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl IPlugin
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterHypothesisServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewHypothesisClient(c)}, nil
}
