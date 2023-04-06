package repo

import "genproto/common"

type CompanyI interface {
	Upsert(*common.CompanyCreatedModel) error
	Delete(*common.RequestID) error
}
