package dag_builder

import "github.com/StandardRunbook/hypothecary/app/config"

type PluginDAG interface {
	StartChain() error
	BuildChain(config *config.Config) error
}
