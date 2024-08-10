package file_permissions_check

import (
	"github.com/StandardRunbook/hypothecary/plugins"
)

type FilePermissionsCheck struct {
}

func NewFilePermissionsCheck() plugins.IPlugin {
	return &FilePermissionsCheck{}
}

func (f *FilePermissionsCheck) Name() string {
	//TODO implement me
	panic("implement me")
}

func (f *FilePermissionsCheck) Run() string {
	//TODO implement me
	panic("implement me")
}

func (f *FilePermissionsCheck) ParseOutput() string {
	//TODO implement me
	panic("implement me")
}
