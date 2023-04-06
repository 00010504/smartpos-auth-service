package listeners

import (
	"context"
	"genproto/common"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/config"
	"github.com/Invan2/invan_auth_service/events/topics"
	"github.com/Invan2/invan_auth_service/pkg/logger"
)

func (a *authService) CreateEmployee(ctx context.Context, req *auth_service.CreateEmployeeRequest) (*common.ResponseID, error) {

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
	res, err := tr.Employee().Create(req)
	if err != nil {
		a.log.Error("error while create user", logger.Error(err))
		return nil, err
	}

	err = a.kafka.Push(topics.UserUpsertTopic, &common.UserCreatedModel{
		Id:          res.Id,
		UserTypeId:  config.SystemUserTypeID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Image:       "",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) GetAllEmployees(ctx context.Context, req *auth_service.GetAllEmployeesRequest) (*auth_service.GetAllEmployeesResponse, error) {

	return a.strg.Employee().GetAll(req)
}

func (a *authService) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if err := tr.Rollback(); err != nil {
				a.log.Error("error while rollback trasaction", logger.Error(err))
			}
		} else {
			if err := tr.Commit(); err != nil {
				a.log.Error("error while commit trasaction", logger.Error(err))
			}
		}
	}()

	res, err := tr.Employee().Delete(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) UpdateEmployee(ctx context.Context, req *auth_service.UpdateEmployeeRequest) (*common.ResponseID, error) {

	tr, err := a.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if err := tr.Rollback(); err != nil {
				a.log.Error("error while rollback trasaction", logger.Error(err))
			}
		} else {
			if err := tr.Commit(); err != nil {
				a.log.Error("error while commit trasaction", logger.Error(err))
			}
		}
	}()

	res, err := tr.Employee().Update(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *authService) GetEmployeeById(ctx context.Context, in *common.RequestID) (*auth_service.ShortEmployee, error) {
	return a.strg.Employee().GetEmployeeById(ctx, in)
}
