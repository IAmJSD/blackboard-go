// Code generated by go-swagger; DO NOT EDIT.

package institutional_hierarchy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDReader is a Reader for the PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeID structure.
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK creates a PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK with default headers values
func NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK() *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK {
	return &PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK{}
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK struct {
	Payload *models.Node
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}][%d] patchLearnApiPublicV1InstitutionalHierarchyNodesNodeIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest creates a PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest() *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest {
	return &PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest{}
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}][%d] patchLearnApiPublicV1InstitutionalHierarchyNodesNodeIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden creates a PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden with default headers values
func NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden() *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden {
	return &PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden{}
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden handles this case with default header values.

Forbidden
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}][%d] patchLearnApiPublicV1InstitutionalHierarchyNodesNodeIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound creates a PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound with default headers values
func NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound() *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound {
	return &PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound{}
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound handles this case with default header values.

Not Found
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}][%d] patchLearnApiPublicV1InstitutionalHierarchyNodesNodeIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict creates a PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict with default headers values
func NewPatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict() *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict {
	return &PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict{}
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict handles this case with default header values.

Conflict
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}][%d] patchLearnApiPublicV1InstitutionalHierarchyNodesNodeIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody patch learn API public v1 institutional hierarchy nodes node ID body
swagger:model PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody
*/
type PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody struct {

	// Node description
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Node unique identifier
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// Node display name
	// Max Length: 256
	Title string `json:"title,omitempty"`
}

// Validate validates this patch learn API public v1 institutional hierarchy nodes node ID body
func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"description", "body", string(o.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"externalId", "body", string(o.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) validateTitle(formats strfmt.Registry) error {

	if swag.IsZero(o.Title) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"title", "body", string(o.Title), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}