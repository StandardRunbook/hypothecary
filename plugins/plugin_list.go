package plugins

import (
	awssts "github.com/StandardRunbook/hypothecary/plugins/aws-get-caller-identity"
	certexpirycheck "github.com/StandardRunbook/hypothecary/plugins/cert-expiry-check"
	filepermissionscheck "github.com/StandardRunbook/hypothecary/plugins/file-permissions-check"
	lsfileexists "github.com/StandardRunbook/hypothecary/plugins/ls-file-exists"
)

type IPlugin interface {
	Name() string
	Run() string
	ParseOutput() string
}

// PluginList is the source of truth for all the plugins
var PluginList = map[string]func() IPlugin{
	"aws-get-caller-identity": awssts.NewGetCallerIdentityAWS,
	"ls-file-exists":          lsfileexists.NewLsFileExists,
	"cert-expiry-check":       certexpirycheck.NewCertExpiryCheck,
	"file-permissions-check":  filepermissionscheck.NewFilePermissionsCheck,
}
