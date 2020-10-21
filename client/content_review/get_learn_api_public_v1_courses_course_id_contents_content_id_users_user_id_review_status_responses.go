// Code generated by go-swagger; DO NOT EDIT.

package content_review

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatus structure.
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK creates a GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK {
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK struct {
	Payload *models.ReviewStatus
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] getLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) GetPayload() *models.ReviewStatus {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReviewStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden creates a GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] getLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound creates a GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] getLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
