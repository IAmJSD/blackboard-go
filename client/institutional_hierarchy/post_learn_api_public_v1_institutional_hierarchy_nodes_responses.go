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

// PostLearnAPIPublicV1InstitutionalHierarchyNodesReader is a Reader for the PostLearnAPIPublicV1InstitutionalHierarchyNodes structure.
type PostLearnAPIPublicV1InstitutionalHierarchyNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1InstitutionalHierarchyNodesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV1InstitutionalHierarchyNodesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesCreated creates a PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated with default headers values
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesCreated() *PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated {
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated{}
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated struct {
	Payload *models.Node
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/institutionalHierarchy/nodes][%d] postLearnApiPublicV1InstitutionalHierarchyNodesCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated) GetPayload() *models.Node {
	return o.Payload
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden creates a PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden with default headers values
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden() *PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden {
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden{}
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden handles this case with default header values.

Forbidden
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/institutionalHierarchy/nodes][%d] postLearnApiPublicV1InstitutionalHierarchyNodesForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesConflict creates a PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict with default headers values
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesConflict() *PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict {
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict{}
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict handles this case with default header values.

Conflict
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/institutionalHierarchy/nodes][%d] postLearnApiPublicV1InstitutionalHierarchyNodesConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesBody post learn API public v1 institutional hierarchy nodes body
swagger:model PostLearnAPIPublicV1InstitutionalHierarchyNodesBody
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesBody struct {

	// Node description
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Node unique identifier
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// Node display name
	// Required: true
	// Max Length: 256
	Title *string `json:"title"`
}

// Validate validates this post learn API public v1 institutional hierarchy nodes body
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) Validate(formats strfmt.Registry) error {
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

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"description", "body", string(o.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"externalId", "body", string(o.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"title", "body", string(*o.Title), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1InstitutionalHierarchyNodesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
