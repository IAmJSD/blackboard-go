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

// GetLearnAPIPublicV2CoursesCourseIDReader is a Reader for the GetLearnAPIPublicV2CoursesCourseID structure.
type GetLearnAPIPublicV2CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV2CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV2CoursesCourseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV2CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV2CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDOK creates a GetLearnAPIPublicV2CoursesCourseIDOK with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDOK() *GetLearnAPIPublicV2CoursesCourseIDOK {
	return &GetLearnAPIPublicV2CoursesCourseIDOK{}
}

/*GetLearnAPIPublicV2CoursesCourseIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV2CoursesCourseIDOK struct {
	Payload *models.CourseV2
}

func (o *GetLearnAPIPublicV2CoursesCourseIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}][%d] getLearnApiPublicV2CoursesCourseIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDOK) GetPayload() *models.CourseV2 {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CourseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDBadRequest creates a GetLearnAPIPublicV2CoursesCourseIDBadRequest with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDBadRequest() *GetLearnAPIPublicV2CoursesCourseIDBadRequest {
	return &GetLearnAPIPublicV2CoursesCourseIDBadRequest{}
}

/*GetLearnAPIPublicV2CoursesCourseIDBadRequest handles this case with default header values.

Failed to create course; or The request did not specify a valid courseId
*/
type GetLearnAPIPublicV2CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}][%d] getLearnApiPublicV2CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDNotFound creates a GetLearnAPIPublicV2CoursesCourseIDNotFound with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDNotFound() *GetLearnAPIPublicV2CoursesCourseIDNotFound {
	return &GetLearnAPIPublicV2CoursesCourseIDNotFound{}
}

/*GetLearnAPIPublicV2CoursesCourseIDNotFound handles this case with default header values.

The course is not found or is unavailable and the user does not have the permission to view unavailable courses
*/
type GetLearnAPIPublicV2CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}][%d] getLearnApiPublicV2CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
