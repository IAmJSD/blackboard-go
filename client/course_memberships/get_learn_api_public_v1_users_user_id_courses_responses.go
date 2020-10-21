// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

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

// GetLearnAPIPublicV1UsersUserIDCoursesReader is a Reader for the GetLearnAPIPublicV1UsersUserIDCourses structure.
type GetLearnAPIPublicV1UsersUserIDCoursesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1UsersUserIDCoursesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1UsersUserIDCoursesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1UsersUserIDCoursesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1UsersUserIDCoursesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1UsersUserIDCoursesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesOK creates a GetLearnAPIPublicV1UsersUserIDCoursesOK with default headers values
func NewGetLearnAPIPublicV1UsersUserIDCoursesOK() *GetLearnAPIPublicV1UsersUserIDCoursesOK {
	return &GetLearnAPIPublicV1UsersUserIDCoursesOK{}
}

/*GetLearnAPIPublicV1UsersUserIDCoursesOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1UsersUserIDCoursesOK struct {
	Payload *GetLearnAPIPublicV1UsersUserIDCoursesOKBody
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/courses][%d] getLearnApiPublicV1UsersUserIdCoursesOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesOK) GetPayload() *GetLearnAPIPublicV1UsersUserIDCoursesOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1UsersUserIDCoursesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesBadRequest creates a GetLearnAPIPublicV1UsersUserIDCoursesBadRequest with default headers values
func NewGetLearnAPIPublicV1UsersUserIDCoursesBadRequest() *GetLearnAPIPublicV1UsersUserIDCoursesBadRequest {
	return &GetLearnAPIPublicV1UsersUserIDCoursesBadRequest{}
}

/*GetLearnAPIPublicV1UsersUserIDCoursesBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1UsersUserIDCoursesBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/courses][%d] getLearnApiPublicV1UsersUserIdCoursesBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesForbidden creates a GetLearnAPIPublicV1UsersUserIDCoursesForbidden with default headers values
func NewGetLearnAPIPublicV1UsersUserIDCoursesForbidden() *GetLearnAPIPublicV1UsersUserIDCoursesForbidden {
	return &GetLearnAPIPublicV1UsersUserIDCoursesForbidden{}
}

/*GetLearnAPIPublicV1UsersUserIDCoursesForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1UsersUserIDCoursesForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/courses][%d] getLearnApiPublicV1UsersUserIdCoursesForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesNotFound creates a GetLearnAPIPublicV1UsersUserIDCoursesNotFound with default headers values
func NewGetLearnAPIPublicV1UsersUserIDCoursesNotFound() *GetLearnAPIPublicV1UsersUserIDCoursesNotFound {
	return &GetLearnAPIPublicV1UsersUserIDCoursesNotFound{}
}

/*GetLearnAPIPublicV1UsersUserIDCoursesNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1UsersUserIDCoursesNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/courses][%d] getLearnApiPublicV1UsersUserIdCoursesNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1UsersUserIDCoursesOKBody get learn API public v1 users user ID courses o k body
swagger:model GetLearnAPIPublicV1UsersUserIDCoursesOKBody
*/
type GetLearnAPIPublicV1UsersUserIDCoursesOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.UserMembership `json:"results"`
}

// Validate validates this get learn API public v1 users user ID courses o k body
func (o *GetLearnAPIPublicV1UsersUserIDCoursesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1UsersUserIDCoursesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1UsersUserIdCoursesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1UsersUserIDCoursesOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1UsersUserIdCoursesOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1UsersUserIdCoursesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1UsersUserIDCoursesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1UsersUserIDCoursesOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1UsersUserIDCoursesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}