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

// CategoryCourseMembership category course membership
//
// swagger:model CategoryCourseMembership
type CategoryCourseMembership struct {

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

	// The course detailed information.
	//
	// **Since**: 3700.5.0
	// Required: true
	Course *CourseV2 `json:"course"`

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

// Validate validates this category course membership
func (m *CategoryCourseMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCourse(formats); err != nil {
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

func (m *CategoryCourseMembership) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.RequiredString("categoryId", "body", string(m.CategoryID)); err != nil {
		return err
	}

	return nil
}

var categoryCourseMembershipTypeCategoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Course","Organization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		categoryCourseMembershipTypeCategoryTypePropEnum = append(categoryCourseMembershipTypeCategoryTypePropEnum, v)
	}
}

const (

	// CategoryCourseMembershipCategoryTypeCourse captures enum value "Course"
	CategoryCourseMembershipCategoryTypeCourse string = "Course"

	// CategoryCourseMembershipCategoryTypeOrganization captures enum value "Organization"
	CategoryCourseMembershipCategoryTypeOrganization string = "Organization"
)

// prop value enum
func (m *CategoryCourseMembership) validateCategoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, categoryCourseMembershipTypeCategoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CategoryCourseMembership) validateCategoryType(formats strfmt.Registry) error {

	if err := validate.RequiredString("categoryType", "body", string(m.CategoryType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryTypeEnum("categoryType", "body", m.CategoryType); err != nil {
		return err
	}

	return nil
}

func (m *CategoryCourseMembership) validateCourse(formats strfmt.Registry) error {

	if err := validate.Required("course", "body", m.Course); err != nil {
		return err
	}

	if m.Course != nil {
		if err := m.Course.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("course")
			}
			return err
		}
	}

	return nil
}

func (m *CategoryCourseMembership) validateCourseID(formats strfmt.Registry) error {

	if err := validate.RequiredString("courseId", "body", string(m.CourseID)); err != nil {
		return err
	}

	return nil
}

func (m *CategoryCourseMembership) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CategoryCourseMembership) validateDataSourceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("dataSourceId", "body", string(m.DataSourceID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CategoryCourseMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryCourseMembership) UnmarshalBinary(b []byte) error {
	var res CategoryCourseMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
