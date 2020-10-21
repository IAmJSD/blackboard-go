// Code generated by go-swagger; DO NOT EDIT.

package content_resources

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

// GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildren structure.
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK creates a GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources/{resourceId}/children][%d] getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest handles this case with default header values.

The specified resource was not a folder; or The children type parameter was not valid. Use File or Folder
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources/{resourceId}/children][%d] getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden creates a GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden handles this case with default header values.

The user does not have permission to view the specified resource
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources/{resourceId}/children][%d] getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound creates a GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound handles this case with default header values.

The specified resource could not be found in the specified course
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources/{resourceId}/children][%d] getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests creates a GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests() *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests handles this case with default header values.

The endpoint is being overloaded with requests
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources/{resourceId}/children][%d] getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody get learn API public v1 courses course ID resources resource ID children o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Resource `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID resources resource ID children o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdResourcesResourceIdChildrenOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDResourcesResourceIDChildrenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
