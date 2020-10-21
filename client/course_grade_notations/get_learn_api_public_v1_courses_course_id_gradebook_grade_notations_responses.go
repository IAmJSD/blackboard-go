// Code generated by go-swagger; DO NOT EDIT.

package course_grade_notations

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

// GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotations structure.
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK creates a GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK() *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/gradeNotations][%d] getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest() *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/gradeNotations][%d] getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden() *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/gradeNotations][%d] getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound() *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/gradeNotations][%d] getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody get learn API public v1 courses course ID gradebook grade notations o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.GradeNotation `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID gradebook grade notations o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookGradeNotationsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
