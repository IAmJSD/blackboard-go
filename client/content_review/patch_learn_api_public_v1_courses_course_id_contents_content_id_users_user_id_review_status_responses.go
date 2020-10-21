// Code generated by go-swagger; DO NOT EDIT.

package content_review

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatus structure.
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK creates a PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK() *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK struct {
	Payload *models.ReviewStatus
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] patchLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) GetPayload() *models.ReviewStatus {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReviewStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden() *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden handles this case with default header values.

Forbidden
*/
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] patchLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound() *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound handles this case with default header values.

Not Found
*/
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/contents/{contentId}/users/{userId}/reviewStatus][%d] patchLearnApiPublicV1CoursesCourseIdContentsContentIdUsersUserIdReviewStatusNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody patch learn API public v1 courses course ID contents content ID users user ID review status body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody struct {

	// The current status of the content's 'reviewed' attribute.
	//
	// **Since**: 3700.16.0
	Reviewed bool `json:"reviewed,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID contents content ID users user ID review status body
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
