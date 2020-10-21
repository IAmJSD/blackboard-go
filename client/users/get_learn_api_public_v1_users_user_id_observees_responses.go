// Code generated by go-swagger; DO NOT EDIT.

package users

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

// GetLearnAPIPublicV1UsersUserIDObserveesReader is a Reader for the GetLearnAPIPublicV1UsersUserIDObservees structure.
type GetLearnAPIPublicV1UsersUserIDObserveesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1UsersUserIDObserveesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1UsersUserIDObserveesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1UsersUserIDObserveesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1UsersUserIDObserveesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1UsersUserIDObserveesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1UsersUserIDObserveesOK creates a GetLearnAPIPublicV1UsersUserIDObserveesOK with default headers values
func NewGetLearnAPIPublicV1UsersUserIDObserveesOK() *GetLearnAPIPublicV1UsersUserIDObserveesOK {
	return &GetLearnAPIPublicV1UsersUserIDObserveesOK{}
}

/*GetLearnAPIPublicV1UsersUserIDObserveesOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1UsersUserIDObserveesOK struct {
	Payload *GetLearnAPIPublicV1UsersUserIDObserveesOKBody
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/observees][%d] getLearnApiPublicV1UsersUserIdObserveesOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesOK) GetPayload() *GetLearnAPIPublicV1UsersUserIDObserveesOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1UsersUserIDObserveesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDObserveesBadRequest creates a GetLearnAPIPublicV1UsersUserIDObserveesBadRequest with default headers values
func NewGetLearnAPIPublicV1UsersUserIDObserveesBadRequest() *GetLearnAPIPublicV1UsersUserIDObserveesBadRequest {
	return &GetLearnAPIPublicV1UsersUserIDObserveesBadRequest{}
}

/*GetLearnAPIPublicV1UsersUserIDObserveesBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1UsersUserIDObserveesBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/observees][%d] getLearnApiPublicV1UsersUserIdObserveesBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDObserveesForbidden creates a GetLearnAPIPublicV1UsersUserIDObserveesForbidden with default headers values
func NewGetLearnAPIPublicV1UsersUserIDObserveesForbidden() *GetLearnAPIPublicV1UsersUserIDObserveesForbidden {
	return &GetLearnAPIPublicV1UsersUserIDObserveesForbidden{}
}

/*GetLearnAPIPublicV1UsersUserIDObserveesForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1UsersUserIDObserveesForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/observees][%d] getLearnApiPublicV1UsersUserIdObserveesForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDObserveesNotFound creates a GetLearnAPIPublicV1UsersUserIDObserveesNotFound with default headers values
func NewGetLearnAPIPublicV1UsersUserIDObserveesNotFound() *GetLearnAPIPublicV1UsersUserIDObserveesNotFound {
	return &GetLearnAPIPublicV1UsersUserIDObserveesNotFound{}
}

/*GetLearnAPIPublicV1UsersUserIDObserveesNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1UsersUserIDObserveesNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}/observees][%d] getLearnApiPublicV1UsersUserIdObserveesNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1UsersUserIDObserveesOKBody get learn API public v1 users user ID observees o k body
swagger:model GetLearnAPIPublicV1UsersUserIDObserveesOKBody
*/
type GetLearnAPIPublicV1UsersUserIDObserveesOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.User `json:"results"`
}

// Validate validates this get learn API public v1 users user ID observees o k body
func (o *GetLearnAPIPublicV1UsersUserIDObserveesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1UsersUserIDObserveesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1UsersUserIdObserveesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1UsersUserIDObserveesOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1UsersUserIdObserveesOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1UsersUserIdObserveesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1UsersUserIDObserveesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1UsersUserIDObserveesOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1UsersUserIDObserveesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
