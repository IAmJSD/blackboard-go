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

// PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserID structure.
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK creates a PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK() *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK struct {
	Payload *models.AttendanceRecord
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) GetPayload() *models.AttendanceRecord {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttendanceRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest creates a PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest() *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest {
	return &PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden() *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden handles this case with default header values.

The user does not have entitlements to update attendance records
*/
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound() *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound handles this case with default header values.

The attendance record is not found
*/
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody patch learn API public v1 courses course ID meetings meeting ID users user ID body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody struct {

	// The primary id of the meeting.
	MeetingID string `json:"meetingId,omitempty"`

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
	// Enum: [Absent Late Present Excused]
	Status string `json:"status,omitempty"`

	// The learn external id of the user.
	UserID string `json:"userId,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID meetings meeting ID users user ID body
func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Absent","Late","Present","Excused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBodyTypeStatusPropEnum = append(patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBodyTypeStatusPropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusAbsent captures enum value "Absent"
	PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusAbsent string = "Absent"

	// PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusLate captures enum value "Late"
	PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusLate string = "Late"

	// PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusPresent captures enum value "Present"
	PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusPresent string = "Present"

	// PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusExcused captures enum value "Excused"
	PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBodyStatusExcused string = "Excused"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("input"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
