// Code generated by go-swagger; DO NOT EDIT.

package courses

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

// GetLearnAPIPublicV1CoursesCourseIDCrossListSetReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDCrossListSet structure.
type GetLearnAPIPublicV1CoursesCourseIDCrossListSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetOK creates a GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetOK() *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK {
	return &GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/crossListSet][%d] getLearnApiPublicV1CoursesCourseIdCrossListSetOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden creates a GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden() *GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/crossListSet][%d] getLearnApiPublicV1CoursesCourseIdCrossListSetForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound creates a GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound() *GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/crossListSet][%d] getLearnApiPublicV1CoursesCourseIdCrossListSetNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody get learn API public v1 courses course ID cross list set o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.CourseChild `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID cross list set o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdCrossListSetOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdCrossListSetOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdCrossListSetOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDCrossListSetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
