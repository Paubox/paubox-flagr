// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EvaluationEntity evaluation entity
//
// swagger:model evaluationEntity
type EvaluationEntity struct {

	// entity context
	EntityContext interface{} `json:"entityContext,omitempty"`

	// entity ID
	EntityID string `json:"entityID,omitempty"`

	// entity type
	EntityType string `json:"entityType,omitempty"`
}

func (m *EvaluationEntity) AddField(key, value string) {
    v, _ := m.EntityContext.(map[string]interface{})
    v[key] = value
    m.EntityContext = v
}

func (m *EvaluationEntity) AddIfMissingField(key, value string) {
    v, _ := m.EntityContext.(map[string]interface{})
	if (v[key] != nil) {
		return
	}
    v[key] = value
    m.EntityContext = v
}


// Validate validates this evaluation entity
func (m *EvaluationEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this evaluation entity based on context it is used
func (m *EvaluationEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationEntity) UnmarshalBinary(b []byte) error {
	var res EvaluationEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
