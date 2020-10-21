// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1DataSourcesDataSourceIDReader is a Reader for the GetLearnAPIPublicV1DataSourcesDataSourceID structure.
type GetLearnAPIPublicV1DataSourcesDataSourceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1DataSourcesDataSourceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1DataSourcesDataSourceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1DataSourcesDataSourceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1DataSourcesDataSourceIDOK creates a GetLearnAPIPublicV1DataSourcesDataSourceIDOK with default headers values
func NewGetLearnAPIPublicV1DataSourcesDataSourceIDOK() *GetLearnAPIPublicV1DataSourcesDataSourceIDOK {
	return &GetLearnAPIPublicV1DataSourcesDataSourceIDOK{}
}

/*GetLearnAPIPublicV1DataSourcesDataSourceIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1DataSourcesDataSourceIDOK struct {
	Payload *models.DataSource
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/dataSources/{dataSourceId}][%d] getLearnApiPublicV1DataSourcesDataSourceIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDOK) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest creates a GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest with default headers values
func NewGetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest() *GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest {
	return &GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest{}
}

/*GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/dataSources/{dataSourceId}][%d] getLearnApiPublicV1DataSourcesDataSourceIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1DataSourcesDataSourceIDForbidden creates a GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden with default headers values
func NewGetLearnAPIPublicV1DataSourcesDataSourceIDForbidden() *GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden {
	return &GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden{}
}

/*GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/dataSources/{dataSourceId}][%d] getLearnApiPublicV1DataSourcesDataSourceIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1DataSourcesDataSourceIDNotFound creates a GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound with default headers values
func NewGetLearnAPIPublicV1DataSourcesDataSourceIDNotFound() *GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound {
	return &GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound{}
}

/*GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/dataSources/{dataSourceId}][%d] getLearnApiPublicV1DataSourcesDataSourceIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1DataSourcesDataSourceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
