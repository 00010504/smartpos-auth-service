package repo

import (
	"context"
	"genproto/common"

	"genproto/auth_service"
)

type UserI interface {
	Create(entity *auth_service.CreateUserRequest) (*common.ResponseID, error)
	GetByPhoneNumber(entity *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error)
	CreateProfile(entity *auth_service.CreateProfileRequest) (*common.ResponseID, string, error)
	GetProfile(entity *common.Request) (*auth_service.Profile, error)
	UpdateProfile(entity *auth_service.UpdateProfileRequest) (*common.ResponseID, error)
	ChangePassword(*auth_service.UpdatePasswordRequest) (*common.ResponseID, error)
	Get(*common.Request) (*auth_service.User, error)
	SearchUserByPhoneNumber(*auth_service.SearchUserByPhoneNumberRequest) (*auth_service.SearchUSerByPhoneNumberResponse, error)
	SetAsVerified(ctx context.Context, in *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error)
	GetUserPermission(ctx context.Context, req *common.Request) (*auth_service.GetAllUserPermissionResponse, error)
}
