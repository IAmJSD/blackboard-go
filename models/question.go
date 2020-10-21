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

// Question question
//
// swagger:model Question
type Question struct {

	// Feedback displayed to students when their submitted response is correct.
	CorrectResponseFeedback string `json:"correctResponseFeedback,omitempty"`

	// The id of the question.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Feedback displayed to students when their submitted response is incorrect.
	IncorrectResponseFeedback string `json:"incorrectResponseFeedback,omitempty"`

	// Text added to the question as a note for the instructor. It is not intended to be displayed to students.
	InstructorNotes string `json:"instructorNotes,omitempty"`

	// The point value for the question.
	Points float64 `json:"points,omitempty"`

	// Position of the Question on the Assessment Canvas.
	Position int32 `json:"position,omitempty"`

	// question handler
	QuestionHandler *QuestionHandler `json:"questionHandler,omitempty"`

	// The main text content for the question. It may include plain and formatted text, and all kinds of content supported by the full text editor.
	// Required: true
	Text *string `json:"text"`

	// The title of the question.
	Title string `json:"title,omitempty"`
}

// Validate validates this question
func (m *Question) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestionHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Question) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateQuestionHandler(formats strfmt.Registry) error {

	if swag.IsZero(m.QuestionHandler) { // not required
		return nil
	}

	if m.QuestionHandler != nil {
		if err := m.QuestionHandler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("questionHandler")
			}
			return err
		}
	}

	return nil
}

func (m *Question) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Question) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Question) UnmarshalBinary(b []byte) error {
	var res Question
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
