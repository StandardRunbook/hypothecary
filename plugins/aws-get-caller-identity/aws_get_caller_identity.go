package aws_get_caller_identity

import "github.com/StandardRunbook/hypothecary/plugins"

type GetCallerIdentityAWS struct {
}

func NewGetCallerIdentityAWS() plugins.IPlugin {
	return &GetCallerIdentityAWS{}
}

func (a *GetCallerIdentityAWS) Name() string {
	return "aws-get-caller-identity"
}

func (a *GetCallerIdentityAWS) Run() string {
	return "ran code"
}

func (a *GetCallerIdentityAWS) ParseOutput() string {
	//TODO implement me
	panic("implement me")
}
