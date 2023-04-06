package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type shopRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewShopRepo(db *sqlx.DB, log logger.Logger) repo.ShopI {
	return &shopRepo{
		db:  db,
		log: log,
	}
}

func (s *shopRepo) Upsert(entity *common.ShopCreatedModel) error {

	query := `
		INSERT INTO
			"shop"
		(
			id,
			name,
			company_id,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT (id) DO
		UPDATE
			SET
			"name" = $2,
			"company_id" = $3
			;
	`

	_, err := s.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert shop")
	}

	return nil
}
