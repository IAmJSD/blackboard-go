// Code generated by go-swagger; DO NOT EDIT.

package deprecated_courses

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

// PatchLearnAPIPublicV1CoursesCourseIDReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseID structure.
type PatchLearnAPIPublicV1CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDOK creates a PatchLearnAPIPublicV1CoursesCourseIDOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDOK() *PatchLearnAPIPublicV1CoursesCourseIDOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDOK struct {
	Payload *models.Course
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}][%d] patchLearnApiPublicV1CoursesCourseIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDOK) GetPayload() *models.Course {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Course)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDBadRequest creates a PatchLearnAPIPublicV1CoursesCourseIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDBadRequest() *PatchLearnAPIPublicV1CoursesCourseIDBadRequest {
	return &PatchLearnAPIPublicV1CoursesCourseIDBadRequest{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDBadRequest handles this case with default header values.

The request did not specify a valid course
*/
type PatchLearnAPIPublicV1CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}][%d] patchLearnApiPublicV1CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDForbidden() *PatchLearnAPIPublicV1CoursesCourseIDForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDForbidden handles this case with default header values.

The user does not have entitlements to modify courses
*/
type PatchLearnAPIPublicV1CoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}][%d] patchLearnApiPublicV1CoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDNotFound() *PatchLearnAPIPublicV1CoursesCourseIDNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDNotFound handles this case with default header values.

The course was not found or is unavailable and the user does not have the permission to view unavailable courses
*/
type PatchLearnAPIPublicV1CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}][%d] patchLearnApiPublicV1CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDConflict creates a PatchLearnAPIPublicV1CoursesCourseIDConflict with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDConflict() *PatchLearnAPIPublicV1CoursesCourseIDConflict {
	return &PatchLearnAPIPublicV1CoursesCourseIDConflict{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDConflict handles this case with default header values.

A course with the same courseId or externalId already exists
*/
type PatchLearnAPIPublicV1CoursesCourseIDConflict struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}][%d] patchLearnApiPublicV1CoursesCourseIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDBody patch learn API public v1 courses course ID body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDBody struct {

	// Whether guests (users with the role guest) are allowed access to the course. Defaults to true.
	AllowGuests bool `json:"allowGuests,omitempty"`

	// availability
	Availability *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability `json:"availability,omitempty"`

	// The ID of the data source associated with this course. This may optionally be the data source's externalId using the syntax "externalId:math101".
	DataSourceID string `json:"dataSourceId,omitempty"`

	// The description of the course.
	Description string `json:"description,omitempty"`

	// enrollment
	Enrollment *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment `json:"enrollment,omitempty"`

	// An optional externally-defined unique ID for the course. Defaults to the courseId.
	//
	// Formerly known as 'batchUid'.
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// locale
	Locale *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale `json:"locale,omitempty"`

	// The name of the course.
	// Max Length: 333
	Name string `json:"name,omitempty"`

	// This status does not affect availability of the course for viewing in any way. readOnly is valid for both Ultra and Classic courses. If an Ultra course is in readOnly mode, updates are not possible. For a Classic course in readOnly mode, updates are still possible (through Web UI but not through REST i.e. closed is enforced for original courses when updated through REST the same way Ultra courses are blocked) but new notifications are not generated.
	//
	// **Deprecated**: since 3400.8.0; use the v2 endpoint's closedComplete property instead
	ReadOnly bool `json:"readOnly,omitempty"`

	// The ID of the term associated to this course. This may optionally be the term's externalId using the syntax "externalId:spring.2016".
	TermID string `json:"termId,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID body
func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocale(formats); err != nil {
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

func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) validateEnrollment(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"externalId", "body", string(o.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) validateLocale(formats strfmt.Registry) error {

	if swag.IsZero(o.Locale) { // not required
		return nil
	}

	if o.Locale != nil {
		if err := o.Locale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "locale")
			}
			return err
		}
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"name", "body", string(o.Name), 333); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability Availability
//
// Settings controlling availability of the course to students.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability
*/
type PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability struct {

	// Whether the course is currently available to students. Instructors can always access the course if they have 'Access unavailable course' entitlement. If set to 'Term', the course's parent term availability settings will be used.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may access the course. |
	// | No | Students may not access the course. |
	// | Disabled | Disabled by the SIS. Students may not access the course.  **Since**: 3100.0.0 |
	// | Term | Availability is inherited from the term settings. Requires a termId be set. |
	//
	// Enum: [Yes No Disabled Term]
	Available string `json:"available,omitempty"`

	// duration
	Duration *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration `json:"duration,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID params body availability
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","Disabled","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum = append(patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableYes string = "Yes"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableNo captures enum value "No"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableNo string = "No"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableDisabled captures enum value "Disabled"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableDisabled string = "Disabled"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableTerm captures enum value "Term"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityAvailableTerm string = "Term"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(o.Duration) { // not required
		return nil
	}

	if o.Duration != nil {
		if err := o.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability" + "." + "duration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration Duration
//
// Settings controlling the length of time the course is available.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration
*/
type PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration struct {

	// The number of days this course can be used. May only be set if availability.duration.type is FixedNumDays.
	DaysOfUse int32 `json:"daysOfUse,omitempty"`

	// The date this course ends. May only be set if availability.duration.type is DateRange.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date this course starts. May only be set if availability.duration.type is DateRange.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// The intended length of the course.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Continuous | Course is active on an ongoing basis. |
	// | DateRange | Course is only intended to be available between specific date ranges |
	// | FixedNumDays | Course is only available for a set number of days |
	// | Term | Course availablity is dictated by its associated term |
	//
	// Enum: [Continuous DateRange FixedNumDays Term]
	Type string `json:"type,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID params body availability duration
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
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

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Continuous","DateRange","FixedNumDays","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum = append(patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeContinuous captures enum value "Continuous"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeContinuous string = "Continuous"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeDateRange captures enum value "DateRange"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeDateRange string = "DateRange"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeFixedNumDays captures enum value "FixedNumDays"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeFixedNumDays string = "FixedNumDays"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeTerm captures enum value "Term"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDurationTypeTerm string = "Term"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"availability"+"."+"duration"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDParamsBodyAvailabilityDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment Enrollment
//
// Settings controlling how students may enroll in the course.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment
*/
type PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment struct {

	// The enrollment access code associated with this course. May only be set if enrollment.type is SelfEnrollment.
	AccessCode string `json:"accessCode,omitempty"`

	// The date on which enrollment in this course ends. May only be set if enrollment.type is SelfEnrollment.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date on which enrollments are allowed for the course. May only be set if enrollment.type is SelfEnrollment.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// Specifies the enrollment options for the course. Defaults to InstructorLed.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | InstructorLed | Enrollment tasks for the course can only performed by the instructor |
	// | SelfEnrollment | Instructors have the ability to enroll users, and students can also enroll themselves in the course |
	// | EmailEnrollment | Instructors have the ability to enroll users, and students can email requests to the instructor for enrollment |
	//
	// Enum: [InstructorLed SelfEnrollment EmailEnrollment]
	Type string `json:"type,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID params body enrollment
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
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

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV1CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InstructorLed","SelfEnrollment","EmailEnrollment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum = append(patchLearnApiPublicV1CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeInstructorLed captures enum value "InstructorLed"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeInstructorLed string = "InstructorLed"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeSelfEnrollment captures enum value "SelfEnrollment"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeSelfEnrollment string = "SelfEnrollment"

	// PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeEmailEnrollment captures enum value "EmailEnrollment"
	PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollmentTypeEmailEnrollment string = "EmailEnrollment"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) validateType(formats strfmt.Registry) error {

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
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDParamsBodyEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale Locale
//
// Settings controlling localization within the course.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale
*/
type PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale struct {

	// Whether students are forced to use the course's specified locale.
	Force bool `json:"force,omitempty"`

	// The locale of this course.
	ID string `json:"id,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID params body locale
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDParamsBodyLocale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}