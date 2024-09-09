package dag_builder

import "github.com/StandardRunbook/hypothecary/plugins"

type Node struct {
	end  bool
	next *Node
	call plugins.IPlugin
}
