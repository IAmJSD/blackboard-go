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

// AttemptV2 attempt v2
//
// swagger:model AttemptV2
type AttemptV2 struct {

	// Return the attempt date associated with this attempt.
	// Required: true
	// Read Only: true
	// Format: date-time
	AttemptDate strfmt.DateTime `json:"attemptDate"`

	// The date on which this attempt was created.
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// The display grade associated with this attempt.
	// Required: true
	DisplayGrade *DisplayGrade `json:"displayGrade"`

	// Whether the score associated with this attempt is ignored when computing the user's grade for the associated grade column.
	Exempt bool `json:"exempt,omitempty"`

	// The instructor feedback associated with this attempt.
	Feedback string `json:"feedback,omitempty"`

	// The group attempt ID associated with this student attempt.
	// Required: true
	// Read Only: true
	GroupAttemptID string `json:"groupAttemptId"`

	// Whether the score associated with this student attempt was overridden from the associated group attempt score. A value is only returned if the attempt grade is currently overridden.
	// Required: true
	// Read Only: true
	GroupOverride bool `json:"groupOverride"`

	// The primary ID for this attempt.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The date when the attempt was modified.
	//
	// **Since**: 3800.0.0
	// Required: true
	// Read Only: true
	// Format: date-time
	Modified strfmt.DateTime `json:"modified"`

	// The instructor notes associated with this attempt.
	Notes string `json:"notes,omitempty"`

	// The reconciliation mode to use when reconciling the attempt grade.  When modifying reconciliationMode, score is also required.  A new score will not be calculated based on the reconciliationMode
	//
	// **Since**: 3700.2.0
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Average |  |
	// | Highest |  |
	// | Lowest |  |
	// | Custom |  |
	//
	// Enum: [Average Highest Lowest Custom]
	ReconciliationMode string `json:"reconciliationMode,omitempty"`

	// The score associated with this attempt.
	Score float64 `json:"score,omitempty"`

	// The status of this attempt.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | NotAttempted | none of the students in a group has submitted an attempt; applies only to group assessments |
	// | Abandoned |   **Deprecated**: Since 9.1 SP8 unsupported status, undetermined behavior if used. |
	// | InProgress | attempt activity has commenced, but has not been submitted for grading |
	// | Suspended |   **Deprecated**: Since 9.1 SP8 unsupported status, undetermined behavior if used. |
	// | Canceled |   **Deprecated**: Since 9.1 SP8 unsupported status, undetermined behavior if used. |
	// | NeedsGrading | attempt has been submitted for grading, but has not been fully graded |
	// | Completed | a grade has been entered for the attempt |
	// | InProgressAgain | attempt has been graded, but more student activity occurred after the grade was entered; applies only to collaborative tools such as discussions |
	// | NeedsGradingAgain | additional student activity occurring after a grade was entered requires that the attempt be regraded; applies only to collaborative tools such as discussions |
	//
	// Enum: [NotAttempted Abandoned InProgress Suspended Canceled NeedsGrading Completed InProgressAgain NeedsGradingAgain]
	Status string `json:"status,omitempty"`

	// The student comments associated with this attempt.
	StudentComments string `json:"studentComments,omitempty"`

	// The student submission text associated with this attempt.
	StudentSubmission string `json:"studentSubmission,omitempty"`

	// The text grade associated with this attempt.
	Text string `json:"text,omitempty"`

	// The user ID associated with this attempt.  Defaults to the authenticated user on create.  Can be specified as a user other than the authenticated user if the authenticated user has the "course.gradebook.MODIFY" entitlement and the authenticated application has the "course.attempt.nonowner.SUBMIT" entitlement.
	// Required: true
	// Read Only: true
	UserID string `json:"userId"`
}

// Validate validates this attempt v2
func (m *AttemptV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttemptDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayGrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupAttemptID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconciliationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttemptV2) validateAttemptDate(formats strfmt.Registry) error {

	if err := validate.Required("attemptDate", "body", strfmt.DateTime(m.AttemptDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("attemptDate", "body", "date-time", m.AttemptDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateDisplayGrade(formats strfmt.Registry) error {

	if err := validate.Required("displayGrade", "body", m.DisplayGrade); err != nil {
		return err
	}

	if m.DisplayGrade != nil {
		if err := m.DisplayGrade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayGrade")
			}
			return err
		}
	}

	return nil
}

func (m *AttemptV2) validateGroupAttemptID(formats strfmt.Registry) error {

	if err := validate.RequiredString("groupAttemptId", "body", string(m.GroupAttemptID)); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateGroupOverride(formats strfmt.Registry) error {

	if err := validate.Required("groupOverride", "body", bool(m.GroupOverride)); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateModified(formats strfmt.Registry) error {

	if err := validate.Required("modified", "body", strfmt.DateTime(m.Modified)); err != nil {
		return err
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

var attemptV2TypeReconciliationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Average","Highest","Lowest","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attemptV2TypeReconciliationModePropEnum = append(attemptV2TypeReconciliationModePropEnum, v)
	}
}

const (

	// AttemptV2ReconciliationModeAverage captures enum value "Average"
	AttemptV2ReconciliationModeAverage string = "Average"

	// AttemptV2ReconciliationModeHighest captures enum value "Highest"
	AttemptV2ReconciliationModeHighest string = "Highest"

	// AttemptV2ReconciliationModeLowest captures enum value "Lowest"
	AttemptV2ReconciliationModeLowest string = "Lowest"

	// AttemptV2ReconciliationModeCustom captures enum value "Custom"
	AttemptV2ReconciliationModeCustom string = "Custom"
)

// prop value enum
func (m *AttemptV2) validateReconciliationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attemptV2TypeReconciliationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttemptV2) validateReconciliationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReconciliationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateReconciliationModeEnum("reconciliationMode", "body", m.ReconciliationMode); err != nil {
		return err
	}

	return nil
}

var attemptV2TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotAttempted","Abandoned","InProgress","Suspended","Canceled","NeedsGrading","Completed","InProgressAgain","NeedsGradingAgain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attemptV2TypeStatusPropEnum = append(attemptV2TypeStatusPropEnum, v)
	}
}

const (

	// AttemptV2StatusNotAttempted captures enum value "NotAttempted"
	AttemptV2StatusNotAttempted string = "NotAttempted"

	// AttemptV2StatusAbandoned captures enum value "Abandoned"
	AttemptV2StatusAbandoned string = "Abandoned"

	// AttemptV2StatusInProgress captures enum value "InProgress"
	AttemptV2StatusInProgress string = "InProgress"

	// AttemptV2StatusSuspended captures enum value "Suspended"
	AttemptV2StatusSuspended string = "Suspended"

	// AttemptV2StatusCanceled captures enum value "Canceled"
	AttemptV2StatusCanceled string = "Canceled"

	// AttemptV2StatusNeedsGrading captures enum value "NeedsGrading"
	AttemptV2StatusNeedsGrading string = "NeedsGrading"

	// AttemptV2StatusCompleted captures enum value "Completed"
	AttemptV2StatusCompleted string = "Completed"

	// AttemptV2StatusInProgressAgain captures enum value "InProgressAgain"
	AttemptV2StatusInProgressAgain string = "InProgressAgain"

	// AttemptV2StatusNeedsGradingAgain captures enum value "NeedsGradingAgain"
	AttemptV2StatusNeedsGradingAgain string = "NeedsGradingAgain"
)

// prop value enum
func (m *AttemptV2) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attemptV2TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttemptV2) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AttemptV2) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("userId", "body", string(m.UserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttemptV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttemptV2) UnmarshalBinary(b []byte) error {
	var res AttemptV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}