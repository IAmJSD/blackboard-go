// Code generated by go-swagger; DO NOT EDIT.

package attendance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsers structure.
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody get learn API public v1 courses course ID meetings meeting ID users o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.AttendanceRecord `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID meetings meeting ID users o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}