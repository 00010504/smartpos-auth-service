package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"
)

func (a *authService) Login(context.Context, *auth_service.LoginRequest) (*common.ResponseID, error) {
	return nil, nil
}
func (a *authService) Verify(context.Context, *auth_service.VerifyRequest) (*common.ResponseID, error) {
	return nil, nil
}

func (a *authService) ForgotPassword(context.Context, *auth_service.ForgotPasswordRequest) (*common.ResponseID, error) {
	return nil, nil
}
