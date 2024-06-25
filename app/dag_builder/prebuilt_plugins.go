package dag_builder

import (
	"fmt"
	"os"
)

const pluginDir = "../../plugins"

type PrebuiltPluginList struct {
	PrebuiltPlugins map[string]IPlugin
}

func NewPrebuiltPluginList() (*PrebuiltPluginList, error) {
	pluginList, err := getFoldersMap(pluginDir)
	if err != nil {
		return nil, err
	}
	return &PrebuiltPluginList{
		PrebuiltPlugins: pluginList,
	}, nil
}

func (pl *PrebuiltPluginList) AddPlugin(plugin IPlugin) error {
	if _, ok := pl.PrebuiltPlugins[plugin.Name()]; ok {
		return fmt.Errorf("plugin %s already exists", plugin.Name())
	}

	pl.PrebuiltPlugins[plugin.Name()] = plugin
	return nil
}

func (pl *PrebuiltPluginList) GetPlugin(name string) IPlugin {
	return pl.PrebuiltPlugins[name]
}

// getFoldersMap returns a map of all folders in the specified directory.
func getFoldersMap(dir string) (map[string]bool, error) {
	foldersMap := make(map[string]bool)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			foldersMap[file.Name()] = true
		}
	}

	return foldersMap, nil
}
