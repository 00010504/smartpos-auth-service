package repo

import (
	"genproto/common"

	"genproto/auth_service"
)

type RoleI interface {
	Create(*auth_service.CreateRoleRequest) (*common.ResponseID, error)
	Delete(*common.RequestID) (*common.ResponseID, error)
	Update(*auth_service.UpdateRoleRequest) (*common.ResponseID, error)
	GetAll(*common.SearchRequest) (*auth_service.GetAllRolesResponse, error)
	Get(req *common.RequestID) (*auth_service.Role, error)
}
