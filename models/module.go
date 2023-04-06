package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"genproto/common"

	"genproto/auth_service"
)

type Module struct {
	ID          string                  `json:"id,omitempty"`
	Name        string                  `json:"name,omitempty"`
	CreatedBy   *common.ShortUser       `json:"created_by,omitempty"`
	Sections    []*auth_service.Section `json:"sections,omitempty"`
	T           PropertyMap             `son:"translation,omitempty"`
	Translation struct {
		Uz string `json:"uz,omitempty"`
		Ru string `json:"ru,omitempty"`
		En string `json:"en,omitempty"`
	} `json:"translation,omitempty"`
}

type PropertyMap map[string]string

func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	stringI, ok := i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]string) failed.")
	}

	var b = map[string]string{}
	for a := range stringI {
		b[a] = a
	}

	*p = b

	return nil
}
