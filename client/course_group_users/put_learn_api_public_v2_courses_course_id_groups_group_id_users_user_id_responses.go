// Code generated by go-swagger; DO NOT EDIT.

package course_group_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDReader is a Reader for the PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserID structure.
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated creates a PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated with default headers values
func NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated() *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated {
	return &PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated{}
}

/*PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated handles this case with default header values.

Created
*/
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated struct {
	Payload *models.GroupMembership
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v2/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV2CoursesCourseIdGroupsGroupIdUsersUserIdCreated  %+v", 201, o.Payload)
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated) GetPayload() *models.GroupMembership {
	return o.Payload
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest creates a PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest with default headers values
func NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest() *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest {
	return &PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest{}
}

/*PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v2/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV2CoursesCourseIdGroupsGroupIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden creates a PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden with default headers values
func NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden() *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden {
	return &PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden{}
}

/*PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v2/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV2CoursesCourseIdGroupsGroupIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound creates a PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound with default headers values
func NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound() *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound {
	return &PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound{}
}

/*PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound handles this case with default header values.

Not Found
*/
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v2/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV2CoursesCourseIdGroupsGroupIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict creates a PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict with default headers values
func NewPutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict() *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict {
	return &PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict{}
}

/*PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict handles this case with default header values.

Not allowed to enroll before enrollment start date; or Not allowed to enroll after enrollment end date; or Maximum number of members exceeded; or Enrolled group has an attempt submitted; or Group has an attempt submitted; or Enrolled group has a grade; or Group has a grade; or The user is already enrolled in the group
*/
type PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v2/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV2CoursesCourseIdGroupsGroupIdUsersUserIdConflict  %+v", 409, o.Payload)
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
