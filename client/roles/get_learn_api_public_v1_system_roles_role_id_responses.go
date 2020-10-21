// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1SystemRolesRoleIDReader is a Reader for the GetLearnAPIPublicV1SystemRolesRoleID structure.
type GetLearnAPIPublicV1SystemRolesRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1SystemRolesRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1SystemRolesRoleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1SystemRolesRoleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1SystemRolesRoleIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1SystemRolesRoleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDOK creates a GetLearnAPIPublicV1SystemRolesRoleIDOK with default headers values
func NewGetLearnAPIPublicV1SystemRolesRoleIDOK() *GetLearnAPIPublicV1SystemRolesRoleIDOK {
	return &GetLearnAPIPublicV1SystemRolesRoleIDOK{}
}

/*GetLearnAPIPublicV1SystemRolesRoleIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1SystemRolesRoleIDOK struct {
	Payload *models.SystemRole
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/systemRoles/{roleId}][%d] getLearnApiPublicV1SystemRolesRoleIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDOK) GetPayload() *models.SystemRole {
	return o.Payload
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDBadRequest creates a GetLearnAPIPublicV1SystemRolesRoleIDBadRequest with default headers values
func NewGetLearnAPIPublicV1SystemRolesRoleIDBadRequest() *GetLearnAPIPublicV1SystemRolesRoleIDBadRequest {
	return &GetLearnAPIPublicV1SystemRolesRoleIDBadRequest{}
}

/*GetLearnAPIPublicV1SystemRolesRoleIDBadRequest handles this case with default header values.

The request did not specify a valid roleId
*/
type GetLearnAPIPublicV1SystemRolesRoleIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/systemRoles/{roleId}][%d] getLearnApiPublicV1SystemRolesRoleIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDForbidden creates a GetLearnAPIPublicV1SystemRolesRoleIDForbidden with default headers values
func NewGetLearnAPIPublicV1SystemRolesRoleIDForbidden() *GetLearnAPIPublicV1SystemRolesRoleIDForbidden {
	return &GetLearnAPIPublicV1SystemRolesRoleIDForbidden{}
}

/*GetLearnAPIPublicV1SystemRolesRoleIDForbidden handles this case with default header values.

Access Denied
*/
type GetLearnAPIPublicV1SystemRolesRoleIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/systemRoles/{roleId}][%d] getLearnApiPublicV1SystemRolesRoleIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDNotFound creates a GetLearnAPIPublicV1SystemRolesRoleIDNotFound with default headers values
func NewGetLearnAPIPublicV1SystemRolesRoleIDNotFound() *GetLearnAPIPublicV1SystemRolesRoleIDNotFound {
	return &GetLearnAPIPublicV1SystemRolesRoleIDNotFound{}
}

/*GetLearnAPIPublicV1SystemRolesRoleIDNotFound handles this case with default header values.

The system role is not found
*/
type GetLearnAPIPublicV1SystemRolesRoleIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/systemRoles/{roleId}][%d] getLearnApiPublicV1SystemRolesRoleIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1SystemRolesRoleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
