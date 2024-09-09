package dag_builder

import (
	"github.com/StandardRunbook/hypothecary/app/config"
	"github.com/StandardRunbook/hypothecary/plugins"
)

type PluginDAG interface {
	StartChain() error
	BuildChain(config *config.Config) error
	AddTask(func() plugins.IPlugin) error
}
