package dag_builder

import (
	"fmt"
	"github.com/StandardRunbook/hypothecary/app/config"
	"github.com/StandardRunbook/hypothecary/plugins"
)

type PluginDag struct {
	plugins map[string]plugins.IPlugin
}

func NewPluginDag() *PluginDag {
	p := &PluginDag{
		plugins: make(map[string]plugins.IPlugin),
	}

	return &PluginDag{}
}

func (d *PluginDag) BuildChain(cfg *config.Config) error {
	// topological sort for plugins to create dag
	pluginList := make(map[string]func() plugins.IPlugin)
	for _, plugin := range cfg.Plugins {
		if plugin.Type == config.PreBuilt {
			if plug, exists := pluginList[plugin.Name]; !exists {
				return fmt.Errorf("plugin not found")
			} else {
				err := d.AddTask(plug)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (d *PluginDag) StartChain() error {
	return nil
}

func (d *PluginDag) AddTask(plugin func() plugins.IPlugin) error {
	return nil
}
