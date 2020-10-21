// Code generated by go-swagger; DO NOT EDIT.

package terms

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

// GetLearnAPIPublicV1TermsReader is a Reader for the GetLearnAPIPublicV1Terms structure.
type GetLearnAPIPublicV1TermsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1TermsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1TermsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1TermsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1TermsOK creates a GetLearnAPIPublicV1TermsOK with default headers values
func NewGetLearnAPIPublicV1TermsOK() *GetLearnAPIPublicV1TermsOK {
	return &GetLearnAPIPublicV1TermsOK{}
}

/*GetLearnAPIPublicV1TermsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1TermsOK struct {
	Payload *GetLearnAPIPublicV1TermsOKBody
}

func (o *GetLearnAPIPublicV1TermsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/terms][%d] getLearnApiPublicV1TermsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1TermsOK) GetPayload() *GetLearnAPIPublicV1TermsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1TermsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1TermsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1TermsForbidden creates a GetLearnAPIPublicV1TermsForbidden with default headers values
func NewGetLearnAPIPublicV1TermsForbidden() *GetLearnAPIPublicV1TermsForbidden {
	return &GetLearnAPIPublicV1TermsForbidden{}
}

/*GetLearnAPIPublicV1TermsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1TermsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1TermsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/terms][%d] getLearnApiPublicV1TermsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1TermsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1TermsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1TermsOKBody get learn API public v1 terms o k body
swagger:model GetLearnAPIPublicV1TermsOKBody
*/
type GetLearnAPIPublicV1TermsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Term `json:"results"`
}

// Validate validates this get learn API public v1 terms o k body
func (o *GetLearnAPIPublicV1TermsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1TermsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1TermsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1TermsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1TermsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1TermsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1TermsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1TermsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1TermsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
