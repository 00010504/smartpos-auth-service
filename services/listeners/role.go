package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"
)

func (a *authService) CreateRole(ctx context.Context, req *auth_service.CreateRoleRequest) (*common.ResponseID, error) {
	return a.strg.Role().Create(req)
}
func (a *authService) DeleteRole(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	return a.strg.Role().Delete(req)
}
func (a *authService) UpdateRole(ctx context.Context, req *auth_service.UpdateRoleRequest) (*common.ResponseID, error) {
	return a.strg.Role().Update(req)
}
func (a *authService) GetAllRoles(ctx context.Context, req *common.SearchRequest) (*auth_service.GetAllRolesResponse, error) {
	return a.strg.Role().GetAll(req)
}

func (a *authService) GetRole(ctx context.Context, req *common.RequestID) (*auth_service.Role, error) {
	return a.strg.Role().Get(req)
}
