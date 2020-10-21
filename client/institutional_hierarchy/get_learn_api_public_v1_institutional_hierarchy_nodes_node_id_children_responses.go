// Code generated by go-swagger; DO NOT EDIT.

package institutional_hierarchy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenReader is a Reader for the GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildren structure.
type GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK creates a GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK with default headers values
func NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK() *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK {
	return &GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK{}
}

/*GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK struct {
	Payload *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/children][%d] getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK) GetPayload() *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden creates a GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden with default headers values
func NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden() *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden {
	return &GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden{}
}

/*GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/children][%d] getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound creates a GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound with default headers values
func NewGetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound() *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound {
	return &GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound{}
}

/*GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/institutionalHierarchy/nodes/{nodeId}/children][%d] getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody get learn API public v1 institutional hierarchy nodes node ID children o k body
swagger:model GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody
*/
type GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Node `json:"results"`
}

// Validate validates this get learn API public v1 institutional hierarchy nodes node ID children o k body
func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1InstitutionalHierarchyNodesNodeIdChildrenOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}