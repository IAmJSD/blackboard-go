// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1UsersUserIDReader is a Reader for the GetLearnAPIPublicV1UsersUserID structure.
type GetLearnAPIPublicV1UsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1UsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1UsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1UsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1UsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1UsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1UsersUserIDOK creates a GetLearnAPIPublicV1UsersUserIDOK with default headers values
func NewGetLearnAPIPublicV1UsersUserIDOK() *GetLearnAPIPublicV1UsersUserIDOK {
	return &GetLearnAPIPublicV1UsersUserIDOK{}
}

/*GetLearnAPIPublicV1UsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1UsersUserIDOK struct {
	Payload *models.User
}

func (o *GetLearnAPIPublicV1UsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}][%d] getLearnApiPublicV1UsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDBadRequest creates a GetLearnAPIPublicV1UsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV1UsersUserIDBadRequest() *GetLearnAPIPublicV1UsersUserIDBadRequest {
	return &GetLearnAPIPublicV1UsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV1UsersUserIDBadRequest handles this case with default header values.

The id is invalid or not properly formatted
*/
type GetLearnAPIPublicV1UsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}][%d] getLearnApiPublicV1UsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDForbidden creates a GetLearnAPIPublicV1UsersUserIDForbidden with default headers values
func NewGetLearnAPIPublicV1UsersUserIDForbidden() *GetLearnAPIPublicV1UsersUserIDForbidden {
	return &GetLearnAPIPublicV1UsersUserIDForbidden{}
}

/*GetLearnAPIPublicV1UsersUserIDForbidden handles this case with default header values.

The user is not authorized to view the specified User object
*/
type GetLearnAPIPublicV1UsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}][%d] getLearnApiPublicV1UsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1UsersUserIDNotFound creates a GetLearnAPIPublicV1UsersUserIDNotFound with default headers values
func NewGetLearnAPIPublicV1UsersUserIDNotFound() *GetLearnAPIPublicV1UsersUserIDNotFound {
	return &GetLearnAPIPublicV1UsersUserIDNotFound{}
}

/*GetLearnAPIPublicV1UsersUserIDNotFound handles this case with default header values.

The user is not found
*/
type GetLearnAPIPublicV1UsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1UsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/users/{userId}][%d] getLearnApiPublicV1UsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1UsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1UsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
