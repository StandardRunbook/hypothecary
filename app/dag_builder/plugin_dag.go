package dag_builder

import "github.com/StandardRunbook/hypothecary/app/config"

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
			d.AddTask()
		}
	}

}

func (d *PluginDag) StartChain() error {

}

func (d *PluginDag) AddTask(plugin IPlugin) error {

}
