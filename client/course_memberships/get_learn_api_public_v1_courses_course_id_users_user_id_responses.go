// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDUsersUserIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDUsersUserID structure.
type GetLearnAPIPublicV1CoursesCourseIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK creates a GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK() *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK struct {
	Payload *models.CourseMembership
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) GetPayload() *models.CourseMembership {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CourseMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest handles this case with default header values.

Invalid data specified in the request
*/
type GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden creates a GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden() *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden handles this case with default header values.

User has insufficient privileges
*/
type GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound creates a GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound() *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound handles this case with default header values.

Membership not found
*/
type GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
