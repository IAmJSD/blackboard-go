// Code generated by go-swagger; DO NOT EDIT.

package course_groups

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

// GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsReader is a Reader for the GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroups structure.
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK struct {
	Payload *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets/{groupId}/groups][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK) GetPayload() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets/{groupId}/groups][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets/{groupId}/groups][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody get learn API public v2 courses course ID groups sets group ID groups o k body
swagger:model GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.GroupV2 `json:"results"`
}

// Validate validates this get learn API public v2 courses course ID groups sets group ID groups o k body
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV2CoursesCourseIdGroupsSetsGroupIdGroupsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV2CoursesCourseIDGroupsSetsGroupIDGroupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
