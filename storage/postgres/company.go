package postgres

import (
	"context"
	"database/sql"
	"genproto/common"

	"github.com/Invan2/invan_auth_service/pkg/logger"

	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type companyRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCompanyRepo(db *sqlx.DB, log logger.Logger) repo.CompanyI {
	return &companyRepo{
		db:  db,
		log: log,
	}
}

func (c *companyRepo) Upsert(entity *common.CompanyCreatedModel) error {

	tr, err := c.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
			return
		}
		_ = tr.Commit()
	}()

	query := `
		INSERT INTO
			"company"
		(
			"id",
			"name",
			"owner",
			"created_by"
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
			"owner" = $3;
	`

	_, err = tr.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.CreatedBy,
		entity.CreatedBy,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert company")
	}

	query = `
		UPDATE 
			"user_profile"
		SET
			"last_company_id" = $1
		WHERE 	
			"version" = (SELECT "last_version" FROM "user" u WHERE u."id"=$2 AND u.deleted_at=0) AND 
			"user_id" = $2
	`

	_, err = tr.Exec(query, entity.Id, entity.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepo) Delete(req *common.RequestID) error {

	query := `
	  	UPDATE
			"company"
	  	SET
			deleted_at = extract(epoch from now())::bigint
	  	WHERE
			id = $1 AND deleted_at = 0
	`

	_, err := c.db.Exec(
		query,
		req.Id,
	)
	if err != nil {
		return errors.Wrap(err, "error while delete company")
	}

	return nil
}
