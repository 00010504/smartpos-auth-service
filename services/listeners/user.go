package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/events/topics"
)

func (a *authService) GetUserByPhoneNumber(ctx context.Context, req *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error) {

	return a.strg.User().GetByPhoneNumber(req)
}

func (a *authService) SearchUserByPhoneNumber(ctx context.Context, req *auth_service.SearchUserByPhoneNumberRequest) (*auth_service.SearchUSerByPhoneNumberResponse, error) {
	return a.strg.User().SearchUserByPhoneNumber(req)
}

func (a *authService) CreateUser(ctx context.Context, req *auth_service.CreateUserRequest) (*common.ResponseID, error) {

	res, err := a.strg.User().Create(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) CreateProfile(ctx context.Context, req *auth_service.CreateProfileRequest) (*common.ResponseID, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	res, phoneNumber, err := tr.User().CreateProfile(req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	err = a.kafka.Push(topics.UserUpsertTopic, &common.UserCreatedModel{
		Id:          req.UserId,
		UserTypeId:  req.UserTypeId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: phoneNumber,
		Image:       "",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) GetUser(ctx context.Context, req *common.Request) (*auth_service.User, error) {
	return a.strg.User().Get(req)
}

func (a *authService) GetProfile(ctx context.Context, req *common.Request) (*auth_service.Profile, error) {
	return a.strg.User().GetProfile(req)
}

func (a *authService) ChangePassword(ctx context.Context, req *auth_service.UpdatePasswordRequest) (*common.ResponseID, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	res, err := tr.User().ChangePassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) ProfileUpdate(ctx context.Context, req *auth_service.UpdateProfileRequest) (*common.ResponseID, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	res, err := tr.User().UpdateProfile(req)
	if err != nil {
		return nil, err
	}
	err = a.kafka.Push(topics.UserUpsertTopic, &common.UserCreatedModel{
		Id:          req.Request.UserId,
		UserTypeId:  req.Request.UserTypeId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: "",
		Image:       "",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) SetAsVerified(ctx context.Context, in *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	res, err := tr.User().SetAsVerified(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) GetUserPermission(ctx context.Context, req *common.Request) (*auth_service.GetAllUserPermissionResponse, error) {
	return a.strg.User().GetUserPermission(ctx, req)
}
