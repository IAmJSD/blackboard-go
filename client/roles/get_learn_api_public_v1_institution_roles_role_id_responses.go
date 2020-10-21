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

// GetLearnAPIPublicV1InstitutionRolesRoleIDReader is a Reader for the GetLearnAPIPublicV1InstitutionRolesRoleID structure.
type GetLearnAPIPublicV1InstitutionRolesRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1InstitutionRolesRoleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1InstitutionRolesRoleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1InstitutionRolesRoleIDOK creates a GetLearnAPIPublicV1InstitutionRolesRoleIDOK with default headers values
func NewGetLearnAPIPublicV1InstitutionRolesRoleIDOK() *GetLearnAPIPublicV1InstitutionRolesRoleIDOK {
	return &GetLearnAPIPublicV1InstitutionRolesRoleIDOK{}
}

/*GetLearnAPIPublicV1InstitutionRolesRoleIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1InstitutionRolesRoleIDOK struct {
	Payload *models.InstitutionRole
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionRoles/{roleId}][%d] getLearnApiPublicV1InstitutionRolesRoleIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDOK) GetPayload() *models.InstitutionRole {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstitutionRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest creates a GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest with default headers values
func NewGetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest() *GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest {
	return &GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest{}
}

/*GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest handles this case with default header values.

The request did not specify a valid roleId
*/
type GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionRoles/{roleId}][%d] getLearnApiPublicV1InstitutionRolesRoleIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1InstitutionRolesRoleIDNotFound creates a GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound with default headers values
func NewGetLearnAPIPublicV1InstitutionRolesRoleIDNotFound() *GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound {
	return &GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound{}
}

/*GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound handles this case with default header values.

The institution role is not found
*/
type GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionRoles/{roleId}][%d] getLearnApiPublicV1InstitutionRolesRoleIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionRolesRoleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
