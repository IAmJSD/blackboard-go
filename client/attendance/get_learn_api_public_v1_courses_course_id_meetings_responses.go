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

// GetLearnAPIPublicV1CoursesCourseIDMeetingsReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDMeetings structure.
type GetLearnAPIPublicV1CoursesCourseIDMeetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsOK creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsOK() *GetLearnAPIPublicV1CoursesCourseIDMeetingsOK {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings][%d] getLearnApiPublicV1CoursesCourseIdMeetingsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest() *GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings][%d] getLearnApiPublicV1CoursesCourseIdMeetingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden() *GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings][%d] getLearnApiPublicV1CoursesCourseIdMeetingsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody get learn API public v1 courses course ID meetings o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.CourseMeeting `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID meetings o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdMeetingsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdMeetingsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDMeetingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
