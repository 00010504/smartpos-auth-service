package storage

import (
	"context"
	"database/sql"

	"github.com/Invan2/invan_auth_service/pkg/logger"

	"github.com/Invan2/invan_auth_service/storage/postgres"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type repos struct {
	userRepo     repo.UserI
	companyRepo  repo.CompanyI
	roleRepo     repo.RoleI
	moduleRepo   repo.ModuleI
	shopRepo     repo.ShopI
	employeeRepo repo.EmployeeI
}

type repoIs interface {
	User() repo.UserI
	Company() repo.CompanyI
	Role() repo.RoleI
	Module() repo.ModuleI
	Employee() repo.EmployeeI
	Shop() repo.ShopI
}

type storage struct {
	db  *sqlx.DB
	log logger.Logger
	repos
}

type storageTr struct {
	tr *sqlx.Tx
	repos
}

type StorageTrI interface {
	Commit() error
	Rollback() error
	repoIs
}

type StoragePG interface {
	WithTransaction(ctx context.Context) (StorageTrI, error)
	repoIs
}

func NewStoragePg(log logger.Logger, db *sqlx.DB) StoragePG {

	return &storage{
		db:  db,
		log: log,
		repos: repos{

			companyRepo:  postgres.NewCompanyRepo(db, log),
			userRepo:     postgres.NewUserRepo(db, log),
			shopRepo:     postgres.NewShopRepo(db, log),
			roleRepo:     postgres.NewRoleRepo(db, log),
			moduleRepo:   postgres.NewModuleRepo(db, log),
			employeeRepo: postgres.NewEmployeeRepo(db, log),
		},
	}
}

func (s *storage) WithTransaction(ctx context.Context) (StorageTrI, error) {

	tr, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &storageTr{
		tr: tr,
		repos: repos{
			userRepo:     postgres.NewUserRepo(tr, s.log),
			employeeRepo: postgres.NewEmployeeRepo(tr, s.log),
		},
	}, nil
}

func (s *storageTr) Commit() error {
	return s.tr.Commit()
}

func (s *storageTr) Rollback() error {
	return s.tr.Rollback()
}

func (s *repos) User() repo.UserI {
	return s.userRepo
}

func (s *repos) Company() repo.CompanyI {
	return s.companyRepo
}

func (s *repos) Role() repo.RoleI {
	return s.roleRepo
}

func (s *repos) Module() repo.ModuleI {
	return s.moduleRepo
}

func (s *repos) Employee() repo.EmployeeI {
	return s.employeeRepo
}

func (s *repos) Shop() repo.ShopI {
	return s.shopRepo
}
