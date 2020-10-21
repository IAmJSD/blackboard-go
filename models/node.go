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

// Node node
//
// swagger:model Node
type Node struct {

	// Node description
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Node unique identifier
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// The primary ID of the Node in the database
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The ID of the Node parent in the database
	// Required: true
	// Read Only: true
	ParentID string `json:"parentId"`

	// Node display name
	// Required: true
	// Max Length: 256
	Title *string `json:"title"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("externalId", "body", string(m.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateParentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("parentId", "body", string(m.ParentID)); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", string(*m.Title), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
