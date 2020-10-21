// Code generated by go-swagger; DO NOT EDIT.

package attendance

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

// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulk structure.
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent handles this case with default header values.

No Content
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent struct {
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/bulk][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkNoContent ", 204)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/bulk][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden handles this case with default header values.

The user does not have entitlements to create attendace records
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/bulk][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody post learn API public v1 courses course ID meetings meeting ID users bulk body
swagger:model PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody struct {

	// The primary id of the meeting.
	// Required: true
	MeetingID *string `json:"meetingId"`

	// The attendance status of the user.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Absent |  |
	// | Late |  |
	// | Present |  |
	// | Excused |  |
	//
	// Required: true
	// Enum: [Absent Late Present Excused]
	Status *string `json:"status"`

	// The learn external id of the user.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this post learn API public v1 courses course ID meetings meeting ID users bulk body
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMeetingID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) validateMeetingID(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"meetingId", "body", o.MeetingID); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Absent","Late","Present","Excused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkBodyTypeStatusPropEnum = append(postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkBodyTypeStatusPropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusAbsent captures enum value "Absent"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusAbsent string = "Absent"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusLate captures enum value "Late"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusLate string = "Late"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusPresent captures enum value "Present"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusPresent string = "Present"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusExcused captures enum value "Excused"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBodyStatusExcused string = "Excused"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBulkBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	// value enum
	if err := o.validateStatusEnum("input"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"userId", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBulkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
