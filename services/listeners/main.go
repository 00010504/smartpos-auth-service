package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/config"
	"github.com/Invan2/invan_auth_service/events"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/services"
	"github.com/Invan2/invan_auth_service/storage"
)

type authService struct {
	log      logger.Logger
	cfg      config.Config
	strg     storage.StoragePG
	strgR    storage.StorageR
	services services.GRPCServices
	kafka    events.PubSubServer
}
type AuthService interface {
	Login(context.Context, *auth_service.LoginRequest) (*common.ResponseID, error)
	Verify(context.Context, *auth_service.VerifyRequest) (*common.ResponseID, error)
	ProfileUpdate(context.Context, *auth_service.UpdateProfileRequest) (*common.ResponseID, error)
	ForgotPassword(context.Context, *auth_service.ForgotPasswordRequest) (*common.ResponseID, error)
	ChangePassword(context.Context, *auth_service.UpdatePasswordRequest) (*common.ResponseID, error)

	CreateModule(context.Context, *auth_service.CreateModuleRequest) (*common.ResponseID, error)
	GetAllModules(context.Context, *auth_service.GetAllModulesRequest) (*auth_service.GetAllModulesResponse, error)

	CreateSection(context.Context, *auth_service.CreateSectionRequest) (*common.ResponseID, error)

	CreateUser(context.Context, *auth_service.CreateUserRequest) (*common.ResponseID, error)
	CreateProfile(context.Context, *auth_service.CreateProfileRequest) (*common.ResponseID, error)
	GetProfile(context.Context, *common.Request) (*auth_service.Profile, error)
	GetUser(context.Context, *common.Request) (*auth_service.User, error)
	GetUserByPhoneNumber(context.Context, *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error)
	SearchUserByPhoneNumber(ctx context.Context, req *auth_service.SearchUserByPhoneNumberRequest) (*auth_service.SearchUSerByPhoneNumberResponse, error)
	GetUserPermission(ctx context.Context, req *common.Request) (*auth_service.GetAllUserPermissionResponse, error)

	CreateRole(context.Context, *auth_service.CreateRoleRequest) (*common.ResponseID, error)
	DeleteRole(context.Context, *common.RequestID) (*common.ResponseID, error)
	UpdateRole(context.Context, *auth_service.UpdateRoleRequest) (*common.ResponseID, error)
	GetAllRoles(context.Context, *common.SearchRequest) (*auth_service.GetAllRolesResponse, error)
	GetRole(context.Context, *common.RequestID) (*auth_service.Role, error)

	CreateEmployee(context.Context, *auth_service.CreateEmployeeRequest) (*common.ResponseID, error)
	GetAllEmployees(context.Context, *auth_service.GetAllEmployeesRequest) (*auth_service.GetAllEmployeesResponse, error)
	Delete(context.Context, *common.RequestID) (*common.ResponseID, error)
	UpdateEmployee(context.Context, *auth_service.UpdateEmployeeRequest) (*common.ResponseID, error)
	GetEmployeeById(ctx context.Context, in *common.RequestID) (*auth_service.ShortEmployee, error)

	SetAsVerified(ctx context.Context, in *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error)
}

func NewAuthService(log logger.Logger, cfg config.Config, services services.GRPCServices, strg storage.StoragePG, strgR storage.StorageR, kafka events.PubSubServer) AuthService {
	return &authService{
		log:      log,
		cfg:      cfg,
		services: services,
		strg:     strg,
		strgR:    strgR,
		kafka:    kafka,
	}
}
