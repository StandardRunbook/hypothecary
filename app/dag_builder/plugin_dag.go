package dag_builder

import (
	"fmt"
	"github.com/StandardRunbook/hypothecary/app/config"
	"github.com/StandardRunbook/hypothecary/plugins"
)

type PluginDag struct {
	plugins map[string]IPlugin
}

func NewPluginDag() *PluginDag {
	return &PluginDag{}
}

func (d *PluginDag) BuildChain(cfg *config.Config) error {
	// topological sort for plugins to create dag
	pluginList := make(map[string]bool)
	for _, plugin := range cfg.Plugins {
		pluginList[plugin.Name] = true
		if plugin.Type == config.PreBuilt {
			if plug, exists := pluginList[plugin.Name]; !exists {
				return fmt.Errorf("plugin not found")
			} else {
				d.AddTask(plug)
			}
		}
	}
	return nil
}

func (d *PluginDag) StartChain() error {
	return nil
}

func (d *PluginDag) AddTask(plugin func() plugins.IPlugin) error {

}
