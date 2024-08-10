package ls_file_exists

import "github.com/StandardRunbook/hypothecary/plugins"

type LsFileExists struct {
}

func NewLsFileExists() plugins.IPlugin {
	return &LsFileExists{}
}

func (l *LsFileExists) Name() string {
	//TODO implement me
	panic("implement me")
}

func (l *LsFileExists) Run() string {
	//TODO implement me
	panic("implement me")
}

func (l *LsFileExists) ParseOutput() string {
	//TODO implement me
	panic("implement me")
}
