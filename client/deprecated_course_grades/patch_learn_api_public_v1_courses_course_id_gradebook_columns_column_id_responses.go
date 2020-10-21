// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_grades

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

// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnID structure.
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK() *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK struct {
	Payload *models.GradeColumn
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK) GetPayload() *models.GradeColumn {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradeColumn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest() *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden() *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden handles this case with default header values.

Forbidden
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound() *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound handles this case with default header values.

Not Found
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody patch learn API public v1 courses course ID gradebook columns column ID body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody struct {

	// availability
	Availability *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability `json:"availability,omitempty"`

	// The description of the grade column.
	Description string `json:"description,omitempty"`

	// Whether this grade column is an external grade column.
	ExternalGrade bool `json:"externalGrade,omitempty"`

	// The externalId for this grade column
	ExternalID string `json:"externalId,omitempty"`

	// grading
	Grading *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading `json:"grading,omitempty"`

	// The name of the grade column.
	Name string `json:"name,omitempty"`

	// score
	Score *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore `json:"score,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook columns column ID body
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) validateGrading(formats strfmt.Registry) error {

	if swag.IsZero(o.Grading) { // not required
		return nil
	}

	if o.Grading != nil {
		if err := o.Grading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "grading")
			}
			return err
		}
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) validateScore(formats strfmt.Registry) error {

	if swag.IsZero(o.Score) { // not required
		return nil
	}

	if o.Score != nil {
		if err := o.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "score")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability Availability
//
// Settings controlling the availability/visibility of grade column data.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability struct {

	// Whether this grade column is available to students
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may view the grade column. |
	// | No | Students may not view the grade column. |
	//
	// Enum: [Yes No]
	Available string `json:"available,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook columns column ID params body availability
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyAvailabilityTypeAvailablePropEnum = append(patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailabilityAvailableYes string = "Yes"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailabilityAvailableNo captures enum value "No"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailabilityAvailableNo string = "No"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

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
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading Grading
//
// Settings controlling whether numerical and text grade values for this grade column are calculated, determined based on attempts, or manually entered.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading struct {

	// anonymous grading
	AnonymousGrading *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading `json:"anonymousGrading,omitempty"`

	// Number of attempts allowed for the grade column.
	AttemptsAllowed int32 `json:"attemptsAllowed,omitempty"`

	// The date on which attempts are due for the grade column.
	// Format: date-time
	Due strfmt.DateTime `json:"due,omitempty"`

	// The scoring model for the submitted grade column attempts.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Last |  |
	// | Highest |  |
	// | Lowest |  |
	// | First |  |
	// | Average |  |
	//
	// Enum: [Last Highest Lowest First Average]
	ScoringModel string `json:"scoringModel,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook columns column ID params body grading
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAnonymousGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDue(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScoringModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) validateAnonymousGrading(formats strfmt.Registry) error {

	if swag.IsZero(o.AnonymousGrading) { // not required
		return nil
	}

	if o.AnonymousGrading != nil {
		if err := o.AnonymousGrading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "grading" + "." + "anonymousGrading")
			}
			return err
		}
	}

	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) validateDue(formats strfmt.Registry) error {

	if swag.IsZero(o.Due) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"grading"+"."+"due", "body", "date-time", o.Due.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingTypeScoringModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Last","Highest","Lowest","First","Average"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingTypeScoringModelPropEnum = append(patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingTypeScoringModelPropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelLast captures enum value "Last"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelLast string = "Last"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelHighest captures enum value "Highest"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelHighest string = "Highest"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelLowest captures enum value "Lowest"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelLowest string = "Lowest"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelFirst captures enum value "First"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelFirst string = "First"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelAverage captures enum value "Average"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingScoringModelAverage string = "Average"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) validateScoringModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingTypeScoringModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) validateScoringModel(formats strfmt.Registry) error {

	if swag.IsZero(o.ScoringModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateScoringModelEnum("input"+"."+"grading"+"."+"scoringModel", "body", o.ScoringModel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading AnonymousGrading
//
// Settings for anonymous grading
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading struct {

	// Date after which grades are released from being anonymized, if AnonymousGrading type is 'Date'.
	// Format: date-time
	ReleaseAfter strfmt.DateTime `json:"releaseAfter,omitempty"`

	// The type of AnonymousGrading settings for this Attempts based Grade Column.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | None | Indicates anonymous grading is not enabled. |
	// | AfterAllGraded | Indicates anonymized grades are released after all attempts have been graded. |
	// | Date | Indicates anonymized grades are released after a specified release date. |
	//
	// Enum: [None AfterAllGraded Date]
	Type string `json:"type,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook columns column ID params body grading anonymous grading
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReleaseAfter(formats); err != nil {
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

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) validateReleaseAfter(formats strfmt.Registry) error {

	if swag.IsZero(o.ReleaseAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"grading"+"."+"anonymousGrading"+"."+"releaseAfter", "body", "date-time", o.ReleaseAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingAnonymousGradingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","AfterAllGraded","Date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingAnonymousGradingTypeTypePropEnum = append(patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingAnonymousGradingTypeTypePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeNone captures enum value "None"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeNone string = "None"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeAfterAllGraded captures enum value "AfterAllGraded"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeAfterAllGraded string = "AfterAllGraded"

	// PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeDate captures enum value "Date"
	PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGradingTypeDate string = "Date"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdParamsBodyGradingAnonymousGradingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"grading"+"."+"anonymousGrading"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyGradingAnonymousGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore Score
//
// Settings controlling the numerical scoring of this grade column.
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore struct {

	// Decimal place precision used to display scores for this grade column.
	//
	// **Deprecated**: since 3200.10.0; no alternative exists since this field never fully functioned as described.
	DecimalPlaces int32 `json:"decimalPlaces,omitempty"`

	// The points possible for this grade column.
	Possible float64 `json:"possible,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook columns column ID params body score
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDParamsBodyScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
