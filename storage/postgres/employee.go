package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"genproto/auth_service"
	"genproto/common"
	"strings"

	"github.com/Invan2/invan_auth_service/pkg/helper"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/lib/pq"

	"github.com/Invan2/invan_auth_service/models"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type employeeRepo struct {
	db  models.DB
	log logger.Logger
}

func NewEmployeeRepo(db models.DB, log logger.Logger) repo.EmployeeI {
	return &employeeRepo{
		db:  db,
		log: log,
	}
}

func (e *employeeRepo) Create(entity *auth_service.CreateEmployeeRequest) (*common.ResponseID, error) {
	var (
		userID     = uuid.NewString()
		shopValues = make([]interface{}, 0)
	)

	query := `
		INSERT INTO "user" ("id", "phone_number", "created_by") VALUES ($1, $2, $3);
	`

	_, err := e.db.Exec(query, userID, entity.PhoneNumber, entity.Request.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "error while create user")
	}

	query = `
		INSERT INTO "user_profile" 
			("id", "user_id", "first_name", "last_name", "password", "last_company_id")
		VALUES
			($1, $2, $3, $4, $5, $6)
	`

	_, err = e.db.Exec(query, uuid.NewString(), userID, entity.FirstName, entity.LastName, entity.Password, entity.Request.CompanyId)
	if err != nil {
		return nil, errors.Wrap(err, "error while create user_profile")
	}

	query = `
		INSERT INTO "user_role" 
			("role_id", "company_id", "user_id", "created_by")
		VALUES
			($1, $2, $3, $4)
`

	_, err = e.db.Exec(query, entity.RoleId, entity.Request.CompanyId, userID, entity.Request.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "error while create user_role")
	}

	query = `
		INSERT INTO "company_user" ("id", "user_id", "company_id", "created_by", "status") VALUES($1, $2, $3, $4, $5);
	`

	// keyinchalik buni olib tashlash krk
	if entity.Status == "" {
		entity.Status = "9fb3ada6-a73b-4b81-9295-5c1605e54552"
	}
	_, err = e.db.Exec(query, uuid.NewString(), userID, entity.Request.CompanyId, entity.Request.UserId, entity.Status)
	if err != nil {
		return nil, errors.Wrap(err, "error while create company user")
	}

	if len(entity.ShopIds) > 0 {

		query = `
			INSERT INTO "shop_user" ( "shop_id","user_id", "created_by") VALUES 
		`

		for _, val := range entity.ShopIds {

			query += `(?, ?, ?),`

			shopValues = append(shopValues, val, userID, entity.Request.UserId)

		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += ` ON CONFLICT ( "shop_id","user_id","deleted_at") DO NOTHING`

		_, err = e.db.Exec(query, shopValues...)
		if err != nil {
			return nil, err
		}

	}

	return &common.ResponseID{Id: userID}, nil
}

func (e *employeeRepo) Update(entity *auth_service.UpdateEmployeeRequest) (*common.ResponseID, error) {

	var (
		shopValues = make([]interface{}, 0)
	)

	query := `UPDATE "company_user" SET status=$1 WHERE user_id=$2 AND company_id=$3 AND deleted_at=0`

	res, err := e.db.Exec(query, entity.Status, entity.Id, entity.Request.CompanyId)
	if err != nil {
		return nil, err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	query = `
		INSERT INTO "shop_user" ( "shop_id","user_id", "created_by") VALUES 
	`

	for _, val := range entity.ShopIds {

		query += `(?, ?, ?),`

		shopValues = append(shopValues, val, entity.Id, entity.Request.UserId)

	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	query += ` ON CONFLICT ( "shop_id","user_id","deleted_at") DO NOTHING`

	_, err = e.db.Exec(query, shopValues...)
	if err != nil {
		return nil, err
	}

	query = `UPDATE shop_user SET deleted_at=extract(epoch from now())::bigint, deleted_by=$3 WHERE user_id=$1 AND (NOT shop_id=ANY($2)) AND deleted_at=0`

	_, err = e.db.Exec(query, entity.Id, pq.Array(entity.ShopIds), entity.Request.UserId)
	if err != nil {
		return nil, err
	}

	query = `INSERT INTO user_role ("role_id", "user_id", "created_by", "company_id") VALUES ($1, $2, $4, $3) ON CONFLICT ("role_id", "user_id", "deleted_at", "company_id") DO NOTHING`

	_, err = e.db.Exec(query, entity.RoleId, entity.Id, entity.Request.CompanyId, entity.Request.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "error while update insert user role")
	}

	query = `UPDATE user_role SET deleted_at=extract(epoch from now())::bigint, deleted_by=$3 WHERE user_id=$1 AND role_id != $2 AND deleted_at=0`

	_, err = e.db.Exec(query, entity.Id, entity.RoleId, entity.Request.UserId)
	if err != nil {
		return nil, err
	}

	return &common.ResponseID{Id: entity.Id}, nil
}

func (e *employeeRepo) GetAll(entity *auth_service.GetAllEmployeesRequest) (*auth_service.GetAllEmployeesResponse, error) {
	var (
		res = auth_service.GetAllEmployeesResponse{
			Employees: make([]*auth_service.ShortEmployee, 0),
		}
		values = map[string]interface{}{
			"limit":      entity.Request.Limit,
			"offset":     entity.Request.Limit * (entity.Request.Page - 1),
			"search":     entity.Request.Search,
			"company_id": entity.Request.Request.CompanyId,
		}
	)

	query := `
		SELECT
			u.id,
			up.first_name,
			up.last_name,
			u.phone_number,
			r.id,
			r.name,
			cu.status,
			us.name,
			us.translation 
		FROM "user" u
		LEFT JOIN "user_profile" up ON up.user_id = u.id AND up.version = u.last_version
		LEFT JOIN "user_role" ur ON ur.user_id = up.user_id AND ur.company_id = :company_id AND ur.deleted_at = 0
		LEFT JOIN "role" r ON r.id = ur.role_id AND r.deleted_at = 0
		JOIN "company_user" cu ON cu.user_id = u.id AND cu.company_id = :company_id AND cu.deleted_at=0
		LEFT JOIN "user_status" us ON us.id=cu.status
	`

	filter := ` WHERE cu.company_id=:company_id AND u.deleted_at = 0 `

	if entity.Request.Search != "" {
		filter += ` AND (
		up."first_name" ILIKE '%' || :search || '%'
		OR up."last_name" ILIKE '%' || :search || '%'
		OR u."phone_number" ILIKE '%' || :search || '%'
		)
		`
	}

	query += filter + `
		ORDER BY  u.created_at
		LIMIT :limit
		OFFSET :offset
	`

	rows, err := e.db.NamedQuery(query, values)
	if err != nil {
		return nil, errors.Wrap(err, "error while searching")
	}

	defer rows.Close()

	for rows.Next() {
		var (
			employe = auth_service.ShortEmployee{
				Status: &common.Status{},
			}
			user models.NullShortUser
			role models.NullRole

			statusTr []byte
		)

		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&role.ID,
			&role.Name,
			&employe.Status.Id,
			&employe.Status.Name,
			&statusTr,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning employes")
		}

		if len(statusTr) > 0 {
			if err := json.Unmarshal(statusTr, &employe.Status.Translation); err != nil {
				return nil, err
			}
		}

		employe.User = &common.ShortUser{
			Id:          user.ID.String,
			FirstName:   user.FirstName.String,
			LastName:    user.LastName.String,
			PhoneNumber: user.PhoneNumber.String,
		}

		employe.Role = &auth_service.ShortRole{
			Id:   role.ID.String,
			Name: role.Name.String,
		}

		res.Employees = append(res.Employees, &employe)
	}

	query = `
		SELECT
			count(u.id)
		from "user" u
		LEFT JOIN "user_profile" up ON up.user_id = u.id AND up.version = u.last_version
		LEFT JOIN "user_role" ur ON ur.user_id = u.id AND ur.company_id = :company_id AND ur.deleted_at = 0
		LEFT JOIN "role" r ON r.id = ur.role_id AND r.deleted_at = 0
		LEFT JOIN "company_user" cu ON cu.user_id = u.id AND cu.company_id = :company_id AND cu.deleted_at = 0
		LEFT JOIN "user_status" us ON us.id=cu.status


	` + filter

	stmt, err := e.db.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrap(err, "error while prepareName")
	}

	defer stmt.Close()

	err = stmt.QueryRow(values).Scan(&res.Total)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning queryRow")
	}

	return &res, nil
}

func (e *employeeRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {

	var userId string

	query := `
		UPDATE
			"company_user"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE
			user_id = $1 AND deleted_at = 0 AND company_id = $2 RETURNING user_id
	`
	err := e.db.QueryRow(query, req.Id, req.Request.CompanyId).Scan(&userId)
	if err != nil {
		return nil, err
	}

	if userId == req.Request.UserId {
		return nil, errors.Wrap(err, "You can not delete self")
	}

	return &common.ResponseID{
		Id: req.Id,
	}, nil
}

func (e *employeeRepo) GetEmployeeById(ctx context.Context, in *common.RequestID) (*auth_service.ShortEmployee, error) {

	var (
		res = auth_service.ShortEmployee{
			User:   &common.ShortUser{},
			Role:   &auth_service.ShortRole{},
			Shops:  make([]*common.ShortShop, 0),
			Status: &common.Status{},
		}

		statusTr []byte
		role     models.NullRole
	)

	query := `
	SELECT
		u.id,
		up.first_name,
		up.last_name,
		u.phone_number,
		r.id,
		r.name,
		cu.status,
		us.name,
		us.translation 
	FROM "user" u
	LEFT JOIN "user_profile" up ON up.user_id = u.id AND up.version = u.last_version
	LEFT JOIN "user_role" ur ON ur.user_id = up.user_id AND ur.company_id = $1 AND ur.deleted_at = 0
	LEFT JOIN "role" r ON r.id = ur.role_id AND r.deleted_at = 0
	JOIN "company_user" cu ON cu.user_id = u.id AND cu.company_id = $1 AND cu.deleted_at=0
	LEFT JOIN "user_status" us ON us.id=cu.status
	WHERE cu.company_id=$1 AND u.deleted_at = 0 AND u.id=$2`

	err := e.db.QueryRow(query, in.Request.CompanyId, in.Id).Scan(&res.User.Id,
		&res.User.FirstName,
		&res.User.LastName,
		&res.User.PhoneNumber,
		&role.ID,
		&role.Name,
		&res.Status.Id,
		&res.Status.Name,
		&statusTr,
	)
	if err != nil {
		return nil, err
	}

	if len(statusTr) > 0 {

		if err := json.Unmarshal(statusTr, &res.Status.Translation); err != nil {
			return nil, err
		}
	}

	query = `
		SELECT
			sh.id,
			sh.name
		FROM shop_user shu
		LEFT JOIN "shop" sh ON sh.id=shop_id AND sh.deleted_at=0
		LEFT JOIN "company" c ON c.id=sh.company_id AND c.deleted_at=0
		WHERE shu.user_id=$1 AND sh.company_id=$2 AND shu.deleted_at=0
	`

	rows, err := e.db.Query(query, in.Id, in.Request.CompanyId)

	defer rows.Close()

	for rows.Next() {
		var (
			shop common.ShortShop
		)

		err := rows.Scan(&shop.Id, &shop.Name)
		if err != nil {
			return nil, err
		}

		res.Shops = append(res.Shops, &shop)
	}

	return &res, nil
}
