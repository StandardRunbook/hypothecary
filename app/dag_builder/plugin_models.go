package dag_builder

import "github.com/StandardRunbook/hypothecary/app/config"

type IPlugin interface {
	Name() string
	Run() string
	ParseOutput() string
}

type PluginDAG interface {
	StartChain() error
	BuildChain(config *config.Config) error
}
