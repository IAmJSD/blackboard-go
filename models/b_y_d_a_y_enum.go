// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BYDAYEnum b y d a y enum
//
// swagger:model BYDAYEnum
type BYDAYEnum string

const (

	// BYDAYEnumSunday captures enum value "Sunday"
	BYDAYEnumSunday BYDAYEnum = "Sunday"

	// BYDAYEnumMonday captures enum value "Monday"
	BYDAYEnumMonday BYDAYEnum = "Monday"

	// BYDAYEnumTuesday captures enum value "Tuesday"
	BYDAYEnumTuesday BYDAYEnum = "Tuesday"

	// BYDAYEnumWednesday captures enum value "Wednesday"
	BYDAYEnumWednesday BYDAYEnum = "Wednesday"

	// BYDAYEnumThursday captures enum value "Thursday"
	BYDAYEnumThursday BYDAYEnum = "Thursday"

	// BYDAYEnumFriday captures enum value "Friday"
	BYDAYEnumFriday BYDAYEnum = "Friday"

	// BYDAYEnumSaturday captures enum value "Saturday"
	BYDAYEnumSaturday BYDAYEnum = "Saturday"
)

// for schema
var bYDAYEnumEnum []interface{}

func init() {
	var res []BYDAYEnum
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bYDAYEnumEnum = append(bYDAYEnumEnum, v)
	}
}

func (m BYDAYEnum) validateBYDAYEnumEnum(path, location string, value BYDAYEnum) error {
	if err := validate.EnumCase(path, location, value, bYDAYEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this b y d a y enum
func (m BYDAYEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBYDAYEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
