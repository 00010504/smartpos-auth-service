package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"
)

func (a *authService) CreateModule(context.Context, *auth_service.CreateModuleRequest) (*common.ResponseID, error) {
	return nil, nil
}
func (a *authService) GetAllModules(ctx context.Context, req *auth_service.GetAllModulesRequest) (*auth_service.GetAllModulesResponse, error) {
	return a.strg.Module().GetAll(req)
}

func (a *authService) CreateSection(context.Context, *auth_service.CreateSectionRequest) (*common.ResponseID, error) {
	return nil, nil
}
