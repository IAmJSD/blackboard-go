// Code generated by go-swagger; DO NOT EDIT.

package deprecated_courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1CoursesCourseIDReader is a Reader for the DeleteLearnAPIPublicV1CoursesCourseID structure.
type DeleteLearnAPIPublicV1CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDNoContent creates a DeleteLearnAPIPublicV1CoursesCourseIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDNoContent() *DeleteLearnAPIPublicV1CoursesCourseIDNoContent {
	return &DeleteLearnAPIPublicV1CoursesCourseIDNoContent{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1CoursesCourseIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}][%d] deleteLearnApiPublicV1CoursesCourseIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDBadRequest creates a DeleteLearnAPIPublicV1CoursesCourseIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDBadRequest() *DeleteLearnAPIPublicV1CoursesCourseIDBadRequest {
	return &DeleteLearnAPIPublicV1CoursesCourseIDBadRequest{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDBadRequest handles this case with default header values.

Invalid courseId provided
*/
type DeleteLearnAPIPublicV1CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}][%d] deleteLearnApiPublicV1CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDForbidden creates a DeleteLearnAPIPublicV1CoursesCourseIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDForbidden() *DeleteLearnAPIPublicV1CoursesCourseIDForbidden {
	return &DeleteLearnAPIPublicV1CoursesCourseIDForbidden{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDForbidden handles this case with default header values.

The user is not authorized to delete the specified Course object
*/
type DeleteLearnAPIPublicV1CoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}][%d] deleteLearnApiPublicV1CoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDNotFound creates a DeleteLearnAPIPublicV1CoursesCourseIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDNotFound() *DeleteLearnAPIPublicV1CoursesCourseIDNotFound {
	return &DeleteLearnAPIPublicV1CoursesCourseIDNotFound{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}][%d] deleteLearnApiPublicV1CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
