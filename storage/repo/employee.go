package repo

import (
	"context"
	"genproto/common"

	"genproto/auth_service"
)

type EmployeeI interface {
	Create(*auth_service.CreateEmployeeRequest) (*common.ResponseID, error)
	GetAll(*auth_service.GetAllEmployeesRequest) (*auth_service.GetAllEmployeesResponse, error)
	Delete(req *common.RequestID) (*common.ResponseID, error)
	Update(entity *auth_service.UpdateEmployeeRequest) (*common.ResponseID, error)
	GetEmployeeById(ctx context.Context, in *common.RequestID) (*auth_service.ShortEmployee, error)
}
