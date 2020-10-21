// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentGroup content group
//
// swagger:model ContentGroup
type ContentGroup struct {

	// The ID of the associated content.
	// Required: true
	// Read Only: true
	ContentID string `json:"contentId"`

	// The ID of the association of content and group.
	// Required: true
	// Read Only: true
	GroupID string `json:"groupId"`
}

// Validate validates this content group
func (m *ContentGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentGroup) validateContentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("contentId", "body", string(m.ContentID)); err != nil {
		return err
	}

	return nil
}

func (m *ContentGroup) validateGroupID(formats strfmt.Registry) error {

	if err := validate.RequiredString("groupId", "body", string(m.GroupID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentGroup) UnmarshalBinary(b []byte) error {
	var res ContentGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
