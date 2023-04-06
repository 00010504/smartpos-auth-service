package postgres

import (
	"context"
	"database/sql"
	"genproto/common"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/models"
	"github.com/Invan2/invan_auth_service/pkg/helper"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type userRepo struct {
	db  models.DB
	log logger.Logger
}

func NewUserRepo(db models.DB, log logger.Logger) repo.UserI {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (u *userRepo) Create(entity *auth_service.CreateUserRequest) (*common.ResponseID, error) {

	id := uuid.NewString()

	query := `
		SELECT 
			u.id
		FROM "user" u
		WHERE u.phone_number = $1 and u.deleted_at = 0 AND u.is_validated;
	`

	err := u.db.QueryRow(query, entity.PhoneNumber).Scan(&id)
	if err != nil && !helper.ErrorIs(err, sql.ErrNoRows.Error()) {
		return nil, err
	}

	if err == nil {
		return nil, errors.New("already exists")
	}

	query = `
		INSERT INTO "user" (
			"id",
			"phone_number"
		) VALUES  (
			$1,
			$2
		);
	`

	_, err = u.db.Exec(query, id, entity.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &common.ResponseID{
		Id: id,
	}, nil
}

func (u *userRepo) GetByPhoneNumber(entity *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error) {

	var res auth_service.GetUserByPhoneNumberResponse
	var (
		profileID, companyID, password sql.NullString
	)

	query := `
		SELECT 
			u.id,
			p.id,
			p.last_company_id,
			p.password,
			u.is_validated
		FROM "user" u
		LEFT JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE u.phone_number = $1 and u.deleted_at = 0;
	`

	err := u.db.QueryRow(query, entity.PhoneNumber).Scan(&res.UserId, &profileID, &companyID, &password, &res.IsVerified)
	if err != nil {
		return nil, errors.Wrap(err, "error while select user profile")
	}

	res.ProfileId = profileID.String
	res.CompanyId = companyID.String
	res.Password = password.String

	return &res, nil
}

func (u *userRepo) SearchUserByPhoneNumber(entity *auth_service.SearchUserByPhoneNumberRequest) (*auth_service.SearchUSerByPhoneNumberResponse, error) {

	var (
		res       auth_service.SearchUSerByPhoneNumberResponse
		companyID string
	)

	query := `
		SELECT
			u.id,
			up.first_name,
			up.last_name,
			cu.company_id
		FROM "user" u
		JOIN "user_profile" up ON up.user_id = u.id AND up.version = u.last_version
		JOIN "company_user" cu ON cu.user_id = u.id AND cu.deleted_at = 0
		WHERE u.phone_number = $1 AND u.deleted_at = 0;
	`

	err := u.db.QueryRow(query, entity.PhoneNumber).Scan(&res.Id, &res.FirstName, &res.LastName, &companyID)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting user")
	}

	res.IsCurrentCompanyUser = companyID == entity.Request.CompanyId

	return &res, nil
}

func (u *userRepo) CreateProfile(entity *auth_service.CreateProfileRequest) (*common.ResponseID, string, error) {

	var (
		phoneNumber string
		version     int
	)

	query := `
		SELECT phone_number, last_version FROM "user" WHERE id = $1 and deleted_at = 0;
	`

	if err := u.db.QueryRow(query, entity.UserId).Scan(&phoneNumber, &version); err != nil {
		return nil, "", err
	}

	query = `
		INSERT INTO "user_profile" 	(
			"id",
			"user_id",
			"first_name",
			"last_name",
			"password",
			"version"
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		);
	`

	id := uuid.NewString()

	version += 1

	_, err := u.db.Exec(query, id, entity.UserId, entity.FirstName, entity.LastName, entity.Password, version)
	if err != nil {
		return nil, "", err
	}

	query = `
		UPDATE "user" SET last_version = $1 WHERE "id" = $2
	`
	_, err = u.db.Exec(query, version, entity.UserId)
	if err != nil {
		return nil, "", err
	}

	return &common.ResponseID{Id: id}, phoneNumber, nil
}

func (u *userRepo) GetProfile(entity *common.Request) (*auth_service.Profile, error) {

	var (
		res   auth_service.Profile
		color sql.NullString
	)

	query := `
		SELECT 
			p.first_name,
			p.last_name,
			p.image,
			p.color
		FROM "user" u
		JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE u.id = $1 AND u.deleted_at = 0
	`

	err := u.db.QueryRow(query, entity.UserId).Scan(&res.FirstName, &res.LastName, &res.Image, &color)
	if err != nil {
		return nil, err
	}

	res.Color = color.String

	return &res, nil
}

func (u *userRepo) UpdateProfile(entity *auth_service.UpdateProfileRequest) (*common.ResponseID, error) {
	var (
		password string
		version  int
	)

	query := `
		SELECT 
			p.password,
			p.version
		FROM "user" u
		JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE u.id = $1 AND u.deleted_at = 0
	`

	if err := u.db.QueryRow(query, entity.Request.UserId).Scan(&password, &version); err != nil {
		return nil, err
	}

	query = `
		INSERT INTO "user_profile" 	(
			"id",
			"user_id",
			"first_name",
			"last_name",
			"password",
			"version",
			"image",
			"color",
			"last_company_id"
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		);
	`

	id := uuid.NewString()
	version += 1

	_, err := u.db.Exec(query, id, entity.Request.UserId, entity.FirstName, entity.LastName, password, version, entity.Image, entity.Color, entity.Request.CompanyId)
	if err != nil {
		return nil, err
	}

	query = `
		UPDATE "user" SET last_version = $1 WHERE "id" = $2
	`
	_, err = u.db.Exec(query, version, entity.Request.UserId)
	if err != nil {
		return nil, err
	}

	return &common.ResponseID{Id: id}, nil
}

func (u *userRepo) ChangePassword(req *auth_service.UpdatePasswordRequest) (*common.ResponseID, error) {

	var (
		user models.User
	)

	query := `
		SELECT 
			u.id,
			u.last_version,
			p.first_name,
			p.last_name,
			COALESCE(p.image, '') as image
		FROM "user" u
		JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE u.id = $1 AND u.deleted_at = 0;
	`

	if err := u.db.QueryRow(query, req.Request.UserId).Scan(&user.ID, &user.LastVersion, &user.FirstName, &user.LastName, &user.Image); err != nil {
		return nil, errors.Wrap(err, "error while scan user")
	}

	query = `
		INSERT INTO
			"user_profile" 
		(
			id,
			user_id,
			first_name,
			last_name,
			password,
			version,
			last_company_id
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		);
	`

	user.LastVersion += 1

	_, err := u.db.Exec(query, uuid.NewString(), user.ID, user.FirstName, user.LastName, req.Password, user.LastVersion, req.Request.CompanyId)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert user_profile")
	}

	query = `
		UPDATE "user" SET last_version = $1 WHERE id = $2
	`
	_, err = u.db.Exec(query, user.LastVersion, user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error while update user last_version")
	}

	return &common.ResponseID{Id: user.ID}, nil
}

func (u *userRepo) Get(entity *common.Request) (*auth_service.User, error) {

	var (
		res auth_service.User
	)

	query := `
		SELECT 
			u.id,
			u.user_type_id,
			p.first_name,
			p.last_name,
			p.password
		FROM "user" u
		JOIN "user_profile" p ON p.user_id = u.id AND p.version = u.last_version
		WHERE u.id = $1 AND u.deleted_at = 0
	`

	err := u.db.QueryRow(query, entity.UserId).Scan(&res.Id, &res.UserTypeId, &res.FirstName, &res.LastName, &res.Password)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *userRepo) SetAsVerified(ctx context.Context, in *auth_service.GetUserByPhoneNumberRequest) (*auth_service.GetUserByPhoneNumberResponse, error) {

	query := `
		UPDATE "user" SET is_validated=true WHERE phone_number=$1
	`

	res, err := u.db.Exec(query, in.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	return u.GetByPhoneNumber(in)
}

func (u *userRepo) GetUserPermission(ctx context.Context, req *common.Request) (*auth_service.GetAllUserPermissionResponse, error) {

	var (
		res = auth_service.GetAllUserPermissionResponse{
			Sections: make([]*auth_service.UserSection, 0),
		}
		sectionIDs    = make([]string, 0)
		permissionMap = make(map[string][]*auth_service.ShortPermission)
	)

	query := `
		SELECT 
			"id",
			"name"
		FROM "section"
	`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var (
			section auth_service.UserSection
		)

		err := rows.Scan(&section.Id, &section.Name)
		if err != nil {
			return nil, err
		}

		sectionIDs = append(sectionIDs, section.Id)
		permissionMap[section.Id] = make([]*auth_service.ShortPermission, 0)

		res.Sections = append(res.Sections, &section)

	}

	rows.Close()

	query = `
		SELECT
			pr.permission_id,
			p.name,
			p."section_id"
		FROM "permission_role" pr
		LEFT JOIN "permission" p ON pr.permission_id = p.id AND p.deleted_at = 0
		WHERE pr.role_id = (SELECT ur.role_id FROM "user_role" ur WHERE ur.user_id = $1 AND ur.company_id = $2 AND ur.deleted_at = 0 AND ur.role_id=pr.role_id) AND pr.deleted_at = 0
		AND p.section_id=ANY($3);
	`

	rows, err = u.db.Query(query, req.UserId, req.CompanyId, pq.Array(sectionIDs))
	if err != nil {
		return nil, err
	}

	sectionIDs = nil

	defer rows.Close()

	for rows.Next() {

		var (
			permission auth_service.ShortPermission
			sectionID  string
		)

		err := rows.Scan(&permission.Id, &permission.Name, &sectionID)
		if err != nil {
			return nil, err
		}

		permissions, ok := permissionMap[sectionID]
		if !ok {
			continue
		}

		permissions = append(permissions, &permission)
		permissionMap[sectionID] = permissions

	}

	for _, section := range res.Sections {
		section.Permissions = permissionMap[section.Id]
	}

	return &res, nil
}
