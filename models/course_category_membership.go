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

// CourseCategoryMembership course category membership
//
// swagger:model CourseCategoryMembership
type CourseCategoryMembership struct {

	// The category detailed information.
	//
	// **Since**: 3700.7.0
	// Required: true
	Category *Category `json:"category"`

	// The category ID.
	// Required: true
	// Read Only: true
	CategoryID string `json:"categoryId"`

	// The category type.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Course |  |
	// | Organization |  |
	//
	// Required: true
	// Read Only: true
	// Enum: [Course Organization]
	CategoryType string `json:"categoryType"`

	// The course ID.
	// Required: true
	// Read Only: true
	CourseID string `json:"courseId"`

	// The date this membership was created.
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// The data source ID.
	// Required: true
	// Read Only: true
	DataSourceID string `json:"dataSourceId"`
}

// Validate validates this course category membership
func (m *CourseCategoryMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCourseID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CourseCategoryMembership) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *CourseCategoryMembership) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.RequiredString("categoryId", "body", string(m.CategoryID)); err != nil {
		return err
	}

	return nil
}

var courseCategoryMembershipTypeCategoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Course","Organization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		courseCategoryMembershipTypeCategoryTypePropEnum = append(courseCategoryMembershipTypeCategoryTypePropEnum, v)
	}
}

const (

	// CourseCategoryMembershipCategoryTypeCourse captures enum value "Course"
	CourseCategoryMembershipCategoryTypeCourse string = "Course"

	// CourseCategoryMembershipCategoryTypeOrganization captures enum value "Organization"
	CourseCategoryMembershipCategoryTypeOrganization string = "Organization"
)

// prop value enum
func (m *CourseCategoryMembership) validateCategoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, courseCategoryMembershipTypeCategoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CourseCategoryMembership) validateCategoryType(formats strfmt.Registry) error {

	if err := validate.RequiredString("categoryType", "body", string(m.CategoryType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryTypeEnum("categoryType", "body", m.CategoryType); err != nil {
		return err
	}

	return nil
}

func (m *CourseCategoryMembership) validateCourseID(formats strfmt.Registry) error {

	if err := validate.RequiredString("courseId", "body", string(m.CourseID)); err != nil {
		return err
	}

	return nil
}

func (m *CourseCategoryMembership) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CourseCategoryMembership) validateDataSourceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("dataSourceId", "body", string(m.DataSourceID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseCategoryMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseCategoryMembership) UnmarshalBinary(b []byte) error {
	var res CourseCategoryMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
