package repo

import "genproto/auth_service"

type ModuleI interface {
	SectionCreate(name string) (string, error)
	GetAll(*auth_service.GetAllModulesRequest) (*auth_service.GetAllModulesResponse, error)
}
