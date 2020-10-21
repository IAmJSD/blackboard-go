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

// GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserID structure.
type GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK() *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden() *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody get learn API public v1 courses course ID meetings users user ID o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.AttendanceRecord `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID meetings users user ID o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
