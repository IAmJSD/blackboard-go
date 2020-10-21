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

// PutLearnAPIPublicV1UsersUserIDObserversObserverIDReader is a Reader for the PutLearnAPIPublicV1UsersUserIDObserversObserverID structure.
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated creates a PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated with default headers values
func NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated() *PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated {
	return &PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated{}
}

/*PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated handles this case with default header values.

Created
*/
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated struct {
	Payload *models.User
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/users/{userId}/observers/{observerId}][%d] putLearnApiPublicV1UsersUserIdObserversObserverIdCreated  %+v", 201, o.Payload)
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest creates a PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest with default headers values
func NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest() *PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest {
	return &PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest{}
}

/*PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest handles this case with default header values.

Bad Request
*/
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/users/{userId}/observers/{observerId}][%d] putLearnApiPublicV1UsersUserIdObserversObserverIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden creates a PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden with default headers values
func NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden() *PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden {
	return &PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden{}
}

/*PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden handles this case with default header values.

Forbidden
*/
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/users/{userId}/observers/{observerId}][%d] putLearnApiPublicV1UsersUserIdObserversObserverIdForbidden  %+v", 403, o.Payload)
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound creates a PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound with default headers values
func NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound() *PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound {
	return &PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound{}
}

/*PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound handles this case with default header values.

Not Found
*/
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/users/{userId}/observers/{observerId}][%d] putLearnApiPublicV1UsersUserIdObserversObserverIdNotFound  %+v", 404, o.Payload)
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict creates a PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict with default headers values
func NewPutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict() *PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict {
	return &PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict{}
}

/*PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict handles this case with default header values.

Conflict
*/
type PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/users/{userId}/observers/{observerId}][%d] putLearnApiPublicV1UsersUserIdObserversObserverIdConflict  %+v", 409, o.Payload)
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1UsersUserIDObserversObserverIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}