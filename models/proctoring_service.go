// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProctoringService proctoring service
//
// swagger:model ProctoringService
type ProctoringService struct {

	// Custom launch parameters for the vendors assessment settings LTI tool.
	AssessmentSettingsCustomParameters map[string]string `json:"assessmentSettingsCustomParameters,omitempty"`

	// The url for the vendors assessment settings LTI tool.
	AssessmentSettingsURL string `json:"assessmentSettingsUrl,omitempty"`

	// availability
	Availability *ProctoringServiceAvailability `json:"availability,omitempty"`

	// The download url for the vendors secure browser.
	BrowserDownloadURL string `json:"browserDownloadUrl,omitempty"`

	// The handle that uniquely identifies this proctoring service.
	// Required: true
	// Read Only: true
	Handle string `json:"handle"`

	// The ID associated with this proctoring service.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The name of the proctoring service.
	// Required: true
	Name *string `json:"name"`

	// The vendors of the proctoring service.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Respondus |  |
	// | Internal |  |
	//
	// Required: true
	// Read Only: true
	// Enum: [Respondus Internal]
	Vendor string `json:"vendor"`
}

// Validate validates this proctoring service
func (m *ProctoringService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProctoringService) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	if m.Availability != nil {
		if err := m.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
			}
			return err
		}
	}

	return nil
}

func (m *ProctoringService) validateHandle(formats strfmt.Registry) error {

	if err := validate.RequiredString("handle", "body", string(m.Handle)); err != nil {
		return err
	}

	return nil
}

func (m *ProctoringService) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ProctoringService) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var proctoringServiceTypeVendorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Respondus","Internal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		proctoringServiceTypeVendorPropEnum = append(proctoringServiceTypeVendorPropEnum, v)
	}
}

const (

	// ProctoringServiceVendorRespondus captures enum value "Respondus"
	ProctoringServiceVendorRespondus string = "Respondus"

	// ProctoringServiceVendorInternal captures enum value "Internal"
	ProctoringServiceVendorInternal string = "Internal"
)

// prop value enum
func (m *ProctoringService) validateVendorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, proctoringServiceTypeVendorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProctoringService) validateVendor(formats strfmt.Registry) error {

	if err := validate.RequiredString("vendor", "body", string(m.Vendor)); err != nil {
		return err
	}

	// value enum
	if err := m.validateVendorEnum("vendor", "body", m.Vendor); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProctoringService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProctoringService) UnmarshalBinary(b []byte) error {
	var res ProctoringService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProctoringServiceAvailability Availability
//
// Settings controlling availability of the proctoring service.
//
// swagger:model ProctoringServiceAvailability
type ProctoringServiceAvailability struct {

	// Whether the proctoring service is available within the system.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes |  |
	// | No |  |
	//
	// Enum: [Yes No]
	Available string `json:"available,omitempty"`
}

// Validate validates this proctoring service availability
func (m *ProctoringServiceAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var proctoringServiceAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		proctoringServiceAvailabilityTypeAvailablePropEnum = append(proctoringServiceAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// ProctoringServiceAvailabilityAvailableYes captures enum value "Yes"
	ProctoringServiceAvailabilityAvailableYes string = "Yes"

	// ProctoringServiceAvailabilityAvailableNo captures enum value "No"
	ProctoringServiceAvailabilityAvailableNo string = "No"
)

// prop value enum
func (m *ProctoringServiceAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, proctoringServiceAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProctoringServiceAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(m.Available) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailableEnum("availability"+"."+"available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProctoringServiceAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProctoringServiceAvailability) UnmarshalBinary(b []byte) error {
	var res ProctoringServiceAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
