package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"genproto/common"
	"strings"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/config"
	"github.com/Invan2/invan_auth_service/models"
	"github.com/Invan2/invan_auth_service/pkg/helper"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type roleRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewRoleRepo(db *sqlx.DB, log logger.Logger) repo.RoleI {
	return &roleRepo{
		db:  db,
		log: log,
	}
}

func (r *roleRepo) Create(entity *auth_service.CreateRoleRequest) (*common.ResponseID, error) {

	tr, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error while start transaction")
	}

	defer func() {

		if err != nil {
			_ = tr.Rollback()
			return
		}

		_ = tr.Commit()

	}()

	roleID := uuid.NewString()
	query := `
		INSERT INTO "role" (
			"id",
			"name",
			"company_id",
			"created_by"
		) VALUES ($1, $2, $3, $4);	
	`
	_, err = tr.Exec(query, roleID, entity.Name, entity.Request.CompanyId, entity.Request.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert role")
	}

	query = `
		INSERT INTO "permission_role" (
			"role_id",
			"permission_id",
			"created_by"
		) VALUES 
	
	`

	values := make([]interface{}, 0)

	for _, permissionID := range entity.PermissionIds {
		values = append(values, roleID, permissionID, entity.Request.UserId)
		query += `(?, ?, ?),`
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := tr.Prepare(query)
	if err != nil {
		return nil, errors.Wrap(err, "error while preparing query")
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert role permissions")
	}

	return &common.ResponseID{Id: roleID}, nil
}

func (r *roleRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {

	var (
		response common.ResponseID
	)

	getUserByRole := `
		SELECT
			user_id
		FROM "user_role"
		WHERE role_id = $1 AND deleted_at = 0
	`
	err := r.db.QueryRow(getUserByRole, req.Id).Scan(&response.Id)
	if err != nil && err.Error() == "no rows in result set" {
		err = nil
	}
	if len(response.Id) > 0 {
		err = errors.New("")
		return nil, errors.Wrap(err, "Cannot delete, because user using this role")
	}

	query := `
		UPDATE "role" SET "deleted_at"=extract(epoch from now())::bigint WHERE id=$1 AND "deleted_at"=0
	`
	res, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	i, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if i == 0 {
		return nil, sql.ErrNoRows
	}

	return &common.ResponseID{
		Id: req.Id,
	}, nil
}

func (r *roleRepo) Update(entity *auth_service.UpdateRoleRequest) (*common.ResponseID, error) {

	var err error

	tr, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error while start transaction")
	}

	defer func() {

		if err != nil {
			_ = tr.Rollback()
			return
		}

		_ = tr.Commit()

	}()

	query := `
	UPDATE 
		"role" 
	SET 
		name = $2
	WHERE 
		id = $1 AND deleted_at = 0 AND company_id = $3
`

	_, err = tr.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error while updating name")
	}

	query = `
		UPDATE
			"permission_role" 
		SET
			deleted_at=extract(epoch from now())::bigint 
		WHERE 
			role_id = $1 AND deleted_at = 0 
	`

	_, err = tr.Exec(
		query,
		entity.Id,
	)
	if err != nil {

		return nil, errors.Wrap(err, "Error while updating role")
	}

	query = `
		INSERT INTO
		"permission_role"
		(
			"role_id",
			"permission_id",
			"created_by"
		) VALUES 
	
	`

	values := make([]interface{}, 0)

	for _, permissionID := range entity.PermissionIds {
		values = append(values, entity.Id, permissionID, entity.Request.UserId)
		query += `(?, ?, ?),`
	}
	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := tr.Prepare(query)
	if err != nil {

		return nil, errors.Wrap(err, "error while preparing query")
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert role permissions")
	}

	return &common.ResponseID{Id: entity.Id}, err
}

func (r *roleRepo) Get(req *common.RequestID) (*auth_service.Role, error) {
	var (
		res = auth_service.Role{
			Modules: make([]*auth_service.Module, 0),
		}
		user models.NullShortUser

		moduleIDs  = make([]string, 0)
		sectionIDs = make([]string, 0)

		sectionsMap = make(map[string][]*auth_service.Section)

		permissionMap = make(map[string][]*auth_service.ShortPermission)
	)

	query := `
		SELECT 
			r."id",
			r."name",
			u.id,
			p.first_name,
			p.last_name,
			(SELECT count(1) FROM "module" m WHERE m.deleted_at=0 AND m.app_type_id=$3) as modules
		FROM "role" r
		LEFT JOIN "user" u ON u.id = r.created_by AND u.deleted_at = 0
		LEFT JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE r.id = $1 AND r.company_id = $2 AND r.deleted_at = 0
	`

	err := r.db.QueryRow(query, req.Id, req.Request.CompanyId, config.SystemTypeID).Scan(&res.Id, &res.Name, &user.ID, &user.FirstName, &user.LastName, &res.NumberOfModules)
	if err != nil {
		return nil, err
	}
	if user.ID.Valid {
		res.CreatedBy = &common.ShortUser{
			Id:        user.ID.String,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
		}
	}

	query = `
		SELECT 
			"id",
			"name",
			"translation"
		FROM "module"
		WHERE "deleted_at" = 0 and app_type_id = $1
	`

	rows, err := r.db.Query(query, config.SystemTypeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			module      auth_service.Module
			translation []byte
		)

		err := rows.Scan(&module.Id, &module.Name, &translation)
		if err != nil {
			return nil, err
		}

		if len(translation) > 0 {
			err := json.Unmarshal(translation, &module.NameTranslation)
			if err != nil {
				return nil, err
			}
		}

		moduleIDs = append(moduleIDs, module.Id)

		sectionsMap[module.Id] = make([]*auth_service.Section, 0)

		res.Modules = append(res.Modules, &module)
	}

	rows.Close()

	query = `
		SELECT 
			"id",
			"name",
			"translation",
			"module_id"
		FROM "section"
		WHERE 
			"module_id"=ANY($1) AND "deleted_at"=0
	`

	rows, err = r.db.Query(query, pq.Array(moduleIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var (
			section     auth_service.Section
			translation []byte
			moduleID    string
		)

		err := rows.Scan(&section.Id, &section.Name, &translation, &moduleID)
		if err != nil {
			return nil, err
		}

		if len(translation) > 0 {
			err := json.Unmarshal(translation, &section.NameTranslation)
			if err != nil {
				return nil, err
			}
		}

		sections, ok := sectionsMap[moduleID]
		if !ok {
			continue
		}
		sections = append(sections, &section)
		sectionIDs = append(sectionIDs, section.Id)
		sectionsMap[moduleID] = sections
		permissionMap[section.Id] = make([]*auth_service.ShortPermission, 0)
	}

	moduleIDs = nil

	rows.Close()

	query = `
		SELECT
			p."id",
			p."name",
			p."section_id",
			p."translation",
			(CASE WHEN pr.role_id IS NULL
				THEN FALSE
				ELSE TRUE
			END) as is_added
		FROM "permission" p
		LEFT JOIN "permission_role" pr ON pr.permission_id=p.id  AND pr.role_id=$2 AND pr.deleted_at=0
		WHERE p.section_id=ANY($1) AND p.deleted_at=0
	`

	rows, err = r.db.Query(query, pq.Array(sectionIDs), res.Id)
	if err != nil {
		return nil, err
	}

	sectionIDs = nil

	defer rows.Close()

	for rows.Next() {
		var (
			permission  auth_service.ShortPermission
			sectionID   string
			translation []byte
		)

		err := rows.Scan(&permission.Id, &permission.Name, &sectionID, &translation, &permission.IsAdded)
		if err != nil {
			return nil, err
		}

		if len(translation) > 0 {
			err := json.Unmarshal(translation, &permission.NameTranslation)
			if err != nil {
				return nil, err
			}
		}

		permissions, ok := permissionMap[sectionID]
		if !ok {
			continue
		}

		permissions = append(permissions, &permission)
		permissionMap[sectionID] = permissions

	}

	for _, module := range res.Modules {

		module.Sections = sectionsMap[module.Id]

		for _, section := range module.Sections {
			section.Permissions = permissionMap[section.Id]
		}

	}

	return &res, nil
}

func (r *roleRepo) GetAll(entity *common.SearchRequest) (*auth_service.GetAllRolesResponse, error) {

	var (
		res auth_service.GetAllRolesResponse

		values = map[string]interface{}{
			"limit":      entity.Limit,
			"offset":     entity.Limit * (entity.Page - 1),
			"search":     entity.Search,
			"company_id": entity.Request.CompanyId,
		}
	)

	query := `
		SELECT
			r."id",
			r."name",
			u.id,
			p.first_name,
			p.last_name,
			r.is_deletable
		FROM "role" r
		LEFT JOIN "user" u ON u.id = r.created_by AND u.deleted_at = 0
		LEFT JOIN "user_profile" p ON p.user_id = u.id and p.version = u.last_version
		WHERE r.company_id=:company_id AND r.deleted_at = 0
	`

	filter := ``

	if entity.Search != "" {
		filter += `AND r."name" ILIKE '%' || :search || '%'`
	}

	query += filter + ` ORDER BY r.created_at `

	query += `LIMIT :limit OFFSET :offset `

	rows, err := r.db.NamedQuery(query, values)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var (
			user models.NullShortUser
			role auth_service.ShortRole
		)

		err := rows.Scan(&role.Id, &role.Name, &user.ID, &user.FirstName, &user.LastName, &role.IsDeletable)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning role")
		}

		if user.ID.Valid {
			role.CreatedBy = &common.ShortUser{
				Id:        user.ID.String,
				FirstName: user.FirstName.String,
				LastName:  user.LastName.String,
			}
		}

		res.Data = append(res.Data, &role)

	}

	query = `
		SELECT
			count(1)
		FROM "role" r
		LEFT JOIN "user" u ON u.id = u.created_by AND u.deleted_at = 0
		LEFT JOIN "user_profile" p ON p.user_id = u.id and p.version = u.last_version
		WHERE r.company_id=:company_id AND r.deleted_at = 0
	` + filter

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(values).Scan(&res.Total)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
