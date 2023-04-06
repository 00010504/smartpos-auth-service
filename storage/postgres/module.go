package postgres

import (
	"encoding/json"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/models"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type moduleRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewModuleRepo(db *sqlx.DB, log logger.Logger) repo.ModuleI {
	return &moduleRepo{
		db:  db,
		log: log,
	}
}

func (p *moduleRepo) SectionCreate(name string) (string, error) {

	var query = `
	INSERT INTO 
		"section"
	(
		id,
		name
	)
	VALUES($1, $2)
	`
	ID := uuid.New().String()
	_, err := p.db.Query(query, ID, name)
	if err != nil {
		return "", errors.Wrap(err, "error while insert section")
	}
	return ID, nil
}

func (p *moduleRepo) SectionUpdate(id string, name string) (string, error) {

	var query = `
	UPDATE
		"section"
	SET 
		name =$2
	WHERE id = $1
	`
	_, err := p.db.Query(query, id, name)
	if err != nil {
		return "", errors.Wrap(err, "error while update section")
	}
	return id, nil
}

func (p *moduleRepo) Get(id string) (*models.Section, error) {

	var section models.Section
	var query = `
	SELECT
		id,
		name,
		created_at
	FROM 
		"section"
	WHERE id = $1
	`

	err := p.db.QueryRow(query, id).Scan(&section.ID, &section.Name, &section.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "error while section get by id")
	}
	return &section, nil

}

func (m *moduleRepo) GetAll(req *auth_service.GetAllModulesRequest) (*auth_service.GetAllModulesResponse, error) {

	var (
		res = auth_service.GetAllModulesResponse{
			Modules: make([]*auth_service.Module, 0),
		}

		moduleIDs  = make([]string, 0)
		sectionIDs = make([]string, 0)

		sectionsMap = make(map[string][]*auth_service.Section)

		permissionMap = make(map[string][]*auth_service.ShortPermission)
	)

	query := `
		SELECT 
			"id",
			"name",
			"translation"
		FROM "module"
		WHERE "deleted_at" = 0 and app_type_id = $1
	`

	rows, err := m.db.Query(query, req.AppType)
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

	rows, err = m.db.Query(query, pq.Array(moduleIDs))
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
			"id",
			"name",
			"section_id",
			"translation"
		FROM "permission"
		WHERE "section_id"=ANY($1) AND "deleted_at"=0
	`

	rows, err = m.db.Query(query, pq.Array(sectionIDs))
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

		err := rows.Scan(&permission.Id, &permission.Name, &sectionID, &translation)
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
