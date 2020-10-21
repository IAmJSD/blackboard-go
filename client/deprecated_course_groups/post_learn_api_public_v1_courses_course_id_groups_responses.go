// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1CoursesCourseIDGroupsReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDGroups structure.
type PostLearnAPIPublicV1CoursesCourseIDGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGroupsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsCreated creates a PostLearnAPIPublicV1CoursesCourseIDGroupsCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsCreated() *PostLearnAPIPublicV1CoursesCourseIDGroupsCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsCreated struct {
	Payload *models.Group
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/groups][%d] postLearnApiPublicV1CoursesCourseIdGroupsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest() *PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/groups][%d] postLearnApiPublicV1CoursesCourseIdGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsForbidden creates a PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsForbidden() *PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden handles this case with default header values.

Forbidden
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/groups][%d] postLearnApiPublicV1CoursesCourseIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsConflict creates a PostLearnAPIPublicV1CoursesCourseIDGroupsConflict with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsConflict() *PostLearnAPIPublicV1CoursesCourseIDGroupsConflict {
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsConflict{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsConflict handles this case with default header values.

Conflict
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/groups][%d] postLearnApiPublicV1CoursesCourseIdGroupsConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsBody post learn API public v1 courses course ID groups body
swagger:model PostLearnAPIPublicV1CoursesCourseIDGroupsBody
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsBody struct {

	// availability
	Availability *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability `json:"availability,omitempty"`

	// The description of the group. This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Description string `json:"description,omitempty"`

	// enrollment
	Enrollment *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment `json:"enrollment,omitempty"`

	// An externally-defined unique ID for the group. Defaults to a random UUID if not provided during create.
	ExternalID string `json:"externalId,omitempty"`

	// Whether or not this is a group set.
	//
	// **Since**: 3700.1.0
	IsGroupSet bool `json:"isGroupSet,omitempty"`

	// The title of the group.
	// Required: true
	Name *string `json:"name"`

	// The primary ID of the group's parent group set.
	ParentID string `json:"parentId,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID groups body
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(o.Availability) { // not required
		return nil
	}

	if o.Availability != nil {
		if err := o.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) validateEnrollment(formats strfmt.Registry) error {

	if swag.IsZero(o.Enrollment) { // not required
		return nil
	}

	if o.Enrollment != nil {
		if err := o.Enrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "enrollment")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGroupsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability Availability
//
// Settings controlling availability of the group to students.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability struct {

	// Whether the Group is currently available to students.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may access the group. |
	// | No | Students may not access the group. |
	// | SignupOnly | Students may only signup and see the group listed, not yet access it. |
	//
	// Enum: [Yes No SignupOnly]
	Available string `json:"available,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID groups params body availability
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","SignupOnly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableNo string = "No"

	// PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableSignupOnly captures enum value "SignupOnly"
	PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailabilityAvailableSignupOnly string = "SignupOnly"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment Enrollment
//
// Settings controlling enrollment of the group to students.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment struct {

	// The maximum allowed enrollment size for self enrolled groups.
	Limit int32 `json:"limit,omitempty"`

	// signup sheet
	SignupSheet *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet `json:"signupSheet,omitempty"`

	// Whether the Group allows self enrollment or only enrolled by instructor. This can only be set on creation.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | InstructorOnly | Students are added to the Group by the instructor |
	// | SelfEnrollment | Students are added to the Group by self enrollment |
	//
	// Enum: [InstructorOnly SelfEnrollment]
	Type string `json:"type,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID groups params body enrollment
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSignupSheet(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) validateSignupSheet(formats strfmt.Registry) error {

	if swag.IsZero(o.SignupSheet) { // not required
		return nil
	}

	if o.SignupSheet != nil {
		if err := o.SignupSheet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "enrollment" + "." + "signupSheet")
			}
			return err
		}
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyEnrollmentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InstructorOnly","SelfEnrollment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyEnrollmentTypeTypePropEnum = append(postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyEnrollmentTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentTypeInstructorOnly captures enum value "InstructorOnly"
	PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentTypeInstructorOnly string = "InstructorOnly"

	// PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentTypeSelfEnrollment captures enum value "SelfEnrollment"
	PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentTypeSelfEnrollment string = "SelfEnrollment"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGroupsParamsBodyEnrollmentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"enrollment"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet SignupSheet
//
// Settings controlling signup to the group for students. Only applicable if Enrollment.Type allows self enrollment.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet struct {

	// The description of the signup sheet This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Description string `json:"description,omitempty"`

	// The name of the signup sheet
	Name string `json:"name,omitempty"`

	// A boolean indicating whether or not members can be seen by others prior to signing up.
	ShowMembers bool `json:"showMembers,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID groups params body enrollment signup sheet
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGroupsParamsBodyEnrollmentSignupSheet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
