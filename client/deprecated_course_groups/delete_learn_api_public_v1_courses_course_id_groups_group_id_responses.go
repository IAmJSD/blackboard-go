// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDReader is a Reader for the DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupID structure.
type DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent creates a DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent() *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/groups/{groupId}][%d] deleteLearnApiPublicV1CoursesCourseIdGroupsGroupIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest creates a DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest() *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/groups/{groupId}][%d] deleteLearnApiPublicV1CoursesCourseIdGroupsGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden creates a DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden() *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/groups/{groupId}][%d] deleteLearnApiPublicV1CoursesCourseIdGroupsGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound creates a DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound() *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/groups/{groupId}][%d] deleteLearnApiPublicV1CoursesCourseIdGroupsGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
