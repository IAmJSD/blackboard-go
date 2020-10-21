// Code generated by go-swagger; DO NOT EDIT.

package institutional_hierarchy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDReader is a Reader for the DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseID structure.
type DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent creates a DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent() *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent {
	return &DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent{}
}

/*DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/courses/{courseId}][%d] deleteLearnApiPublicV1InstitutionalHierarchyNodesNodeIdCoursesCourseIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden creates a DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden() *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden {
	return &DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden{}
}

/*DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/courses/{courseId}][%d] deleteLearnApiPublicV1InstitutionalHierarchyNodesNodeIdCoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound creates a DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound() *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound {
	return &DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound{}
}

/*DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/courses/{courseId}][%d] deleteLearnApiPublicV1InstitutionalHierarchyNodesNodeIdCoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDCoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
