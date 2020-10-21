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

// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDGradebookColumns structure.
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated creates a PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated() *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated struct {
	Payload *models.GradeColumn
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/columns][%d] postLearnApiPublicV1CoursesCourseIdGradebookColumnsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated) GetPayload() *models.GradeColumn {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradeColumn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest() *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/columns][%d] postLearnApiPublicV1CoursesCourseIdGradebookColumnsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict creates a PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict() *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict handles this case with default header values.

Conflict
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/columns][%d] postLearnApiPublicV1CoursesCourseIdGradebookColumnsConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody post learn API public v1 courses course ID gradebook columns body
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody struct {

	// availability
	Availability *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability `json:"availability,omitempty"`

	// The description of the grade column.
	Description string `json:"description,omitempty"`

	// Whether this grade column is an external grade column.
	ExternalGrade bool `json:"externalGrade,omitempty"`

	// The externalId for this grade column
	ExternalID string `json:"externalId,omitempty"`

	// grading
	// Required: true
	Grading *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading `json:"grading"`

	// The name of the grade column.
	// Required: true
	Name *string `json:"name"`

	// score
	Score *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore `json:"score,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID gradebook columns body
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
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

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) validateGrading(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"grading", "body", o.Grading); err != nil {
		return err
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

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) validateScore(formats strfmt.Registry) error {

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
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability Availability
//
// Settings controlling the availability/visibility of grade column data.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability struct {

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

// Validate validates this post learn API public v1 courses course ID gradebook columns params body availability
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailabilityAvailableNo string = "No"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

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
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading Grading
//
// Settings controlling whether numerical and text grade values for this grade column are calculated, determined based on attempts, or manually entered.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading struct {

	// anonymous grading
	// Required: true
	AnonymousGrading *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading `json:"anonymousGrading"`

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

	// The type of Grading settings for this Grade Column.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Attempts | Indicates score and grade values are determined based on user attempts |
	// | Calculated | Indicates score and grade values are determined by applying a calculated formula. |
	// | Manual | Indicates score and grade values are manually entered. |
	//
	// Required: true
	// Enum: [Attempts Calculated Manual]
	Type *string `json:"type"`
}

// Validate validates this post learn API public v1 courses course ID gradebook columns params body grading
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) Validate(formats strfmt.Registry) error {
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

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateAnonymousGrading(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"grading"+"."+"anonymousGrading", "body", o.AnonymousGrading); err != nil {
		return err
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

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateDue(formats strfmt.Registry) error {

	if swag.IsZero(o.Due) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"grading"+"."+"due", "body", "date-time", o.Due.String(), formats); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeScoringModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Last","Highest","Lowest","First","Average"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeScoringModelPropEnum = append(postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeScoringModelPropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelLast captures enum value "Last"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelLast string = "Last"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelHighest captures enum value "Highest"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelHighest string = "Highest"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelLowest captures enum value "Lowest"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelLowest string = "Lowest"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelFirst captures enum value "First"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelFirst string = "First"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelAverage captures enum value "Average"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingScoringModelAverage string = "Average"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateScoringModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeScoringModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateScoringModel(formats strfmt.Registry) error {

	if swag.IsZero(o.ScoringModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateScoringModelEnum("input"+"."+"grading"+"."+"scoringModel", "body", o.ScoringModel); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attempts","Calculated","Manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeTypePropEnum = append(postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeAttempts captures enum value "Attempts"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeAttempts string = "Attempts"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeCalculated captures enum value "Calculated"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeCalculated string = "Calculated"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeManual captures enum value "Manual"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingTypeManual string = "Manual"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) validateType(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"grading"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"grading"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading AnonymousGrading
//
// Settings for anonymous grading
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading struct {

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
	// Required: true
	// Enum: [None AfterAllGraded Date]
	Type *string `json:"type"`
}

// Validate validates this post learn API public v1 courses course ID gradebook columns params body grading anonymous grading
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) Validate(formats strfmt.Registry) error {
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

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) validateReleaseAfter(formats strfmt.Registry) error {

	if swag.IsZero(o.ReleaseAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"grading"+"."+"anonymousGrading"+"."+"releaseAfter", "body", "date-time", o.ReleaseAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingAnonymousGradingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","AfterAllGraded","Date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingAnonymousGradingTypeTypePropEnum = append(postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingAnonymousGradingTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeNone captures enum value "None"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeNone string = "None"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeAfterAllGraded captures enum value "AfterAllGraded"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeAfterAllGraded string = "AfterAllGraded"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeDate captures enum value "Date"
	PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGradingTypeDate string = "Date"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGradebookColumnsParamsBodyGradingAnonymousGradingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) validateType(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"grading"+"."+"anonymousGrading"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"grading"+"."+"anonymousGrading"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyGradingAnonymousGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore Score
//
// Settings controlling the numerical scoring of this grade column.
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore struct {

	// Decimal place precision used to display scores for this grade column.
	//
	// **Deprecated**: since 3200.10.0; no alternative exists since this field never fully functioned as described.
	DecimalPlaces int32 `json:"decimalPlaces,omitempty"`

	// The points possible for this grade column.
	Possible float64 `json:"possible,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID gradebook columns params body score
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsBodyScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}